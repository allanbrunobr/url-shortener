package models

import "time"

type URL struct {
	ID             string    `bson:"_id,omitempty" json:"id"`
	OriginalURL    string    `bson:"original_url" json:"original_url"`
	ShortURL       string    `bson:"short_url" json:"short_url"`
	CreationDate   time.Time `bson:"creation_date,omitempty" json:"creation_date"`
	ExpirationDate time.Time `bson:"expiration_date,omitempty" json:"expiration_date"`
	UserID         int       `bson:"user_id" json:"user_id"`
	ClickCount     int       `bson:"click_count,omitempty" json:"click_count"`
}
