package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/skip2/go-qrcode"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"shorten-url-back-go/models"
	"sync"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/time/rate"
)

var client *mongo.Client

var (
	limiter = rate.NewLimiter(1, 3)
	m       sync.Mutex
)

func main() {

	//connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)

	//configure the router
	r := mux.NewRouter()
	r.Use(rateLimitMiddleware)
	r.HandleFunc("/shorten", ShortenURL).Methods("POST")
	r.HandleFunc("/{shortURL}", RedirectURL).Methods("GET")

	// CORS configuration
	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)

	// start the server with CORS support
	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsOptions(r)))

}

func rateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.Lock()
		defer m.Unlock()
		if !limiter.Allow() {
			http.Error(w, "Muitas requisições, por favor, tente mais tarde", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func validateURL(inputURL string) error {
	_, err := url.ParseRequestURI(inputURL)
	if err != nil {
		return errors.New("Invalid URL")
	}
	return nil
}

// generateShortURL generates a random 6-character string using lowercase and uppercase letters,
// as well as digits. This string will be used as a shortened URL.
//
// The function uses the 'rand' package to generate a random sequence of bytes,
// which is then converted to a string using the provided set of characters.
//
// The generated shortened URL is returned as a string.
func generateShortURL() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// generateQRCode generates a QR code image for the given URL.
//
// The function uses the 'qrcode' package to encode the provided URL into a QR code image.
// The QR code is generated with a medium error correction level and a size of 256x256 pixels.
//
// The function returns the QR code image as a slice of bytes, and any encountered errors.
// If an error occurs during the encoding process, the function returns nil for the QR code image and the error as the second return value.
//
// Otherwise, the function returns the QR code image as a slice of bytes and nil for the error.
func generateQRCode(url string) ([]byte, error) {
	qrCode, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}
	return qrCode, nil
}

// ShortenURL receives a POST request with a URL to be shortened, creates a new
// shortened URL, stores it in the database, and returns the newly created
// shortened URL in the response. It takes a http.ResponseWriter and an
// *http.Request as input parameters.
//
// The function uses the 'json' package to decode the incoming JSON payload
// into a models.URL struct. It then generates a new shortened URL using the
// 'generateShortURL' function, populates the URL struct with the new shortened
// URL and the current time, and inserts the struct into the MongoDB database.
//
// If an error occurs during the insertion process, the function returns an HTTP
// error with a status code of 500 (Internal Server Error) and the error message.
//
// Otherwise, the function returns the newly created shortened URL in the response
// with a status code of 200 (OK).
func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		OriginalURL string `json:"original_url"`
		CustomSlug  string `json:"custom_slug"`
	}

	_ = json.NewDecoder(r.Body).Decode(&requestData)

	err := validateURL(requestData.OriginalURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := client.Database("shortner-url").Collection("url_mapping")
	if requestData.CustomSlug != "" {
		var existingURL models.URL
		err := collection.FindOne(context.TODO(), bson.M{"short_url": requestData.CustomSlug}).Decode(&existingURL)
		if err == nil {
			http.Error(w, "Custom slug is already in use, please choose another one.", http.StatusConflict)
			return
		}
	}

	slug := requestData.CustomSlug
	if slug == "" {
		slug = generateShortURL()
	}

	url := models.URL{
		OriginalURL:  requestData.OriginalURL,
		ShortURL:     slug,
		CreationDate: time.Now(),
		ClickCount:   0,
	}

	qrCode, err := generateQRCode("http://localhost:8080/" + url.ShortURL) // Ajuste aqui
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = collection.InsertOne(context.TODO(), url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"short_url":     "http://localhost:8080/" + url.ShortURL, // Ajuste aqui
		"qr_code":       qrCode,
		"click_count":   url.ClickCount,
		"creation_date": url.CreationDate,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

// RedirectURL handles a GET request with a shortened URL parameter, looks up the
// original URL in the database, increments its click count, and redirects the
// client to the original URL. It takes a http.ResponseWriter and an
// *http.Request as input parameters.
//
// The function extracts the shortened URL parameter from the request's URL
// path, looks up the corresponding original URL in the MongoDB database,
// increments its click count using an update operation, and then redirects the
// client to the original URL using the 'http.Redirect' function.
//
// If an error occurs during the lookup or update process, the function returns an
// HTTP error with a status code of 404 (Not Found) and the error message.
//
// Otherwise, the function redirects the client to the original URL with a status
// code of 302 (Found).
func RedirectURL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	shortURL := params["shortURL"]

	collection := client.Database("shortner-url").Collection("url_mapping")
	var url models.URL
	err := collection.FindOne(context.TODO(), bson.M{"short_url": shortURL}).Decode(&url)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	_, err = collection.UpdateOne(
		context.TODO(),
		bson.M{"short_url": shortURL},
		bson.M{"$inc": bson.M{"click_count": 1}},
	)
	if err != nil {
		log.Printf("Erro ao incrementar o contador de cliques: %v", err)
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusMovedPermanently)
}
