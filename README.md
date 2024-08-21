# URL Shortener

This project is a fullstack application that combines a **React** frontend and a **Golang** backend. It offers a user interface for shortening URLs and a backend server that handles the generation, storage, and redirection of shortened URLs.

## Table of Contents

1. [About the Project](#about-the-project)
2. [Technologies Used](#technologies-used)
3. [Installation and Setup](#installation-and-setup)
   - [Prerequisites](#prerequisites)
   - [Backend Setup](#backend-setup)
   - [Frontend Setup](#frontend-setup)
4. [How to Use](#how-to-use)
5. [Folder Structure](#folder-structure)
6. [Contributing](#contributing)
7. [License](#license)
8. [Contact](#contact)

## About the Project

The **url-shortener** is a URL shortening application that allows users to shorten long URLs and customize the slug of the shortened URL. The frontend provides an intuitive interface for users to interact with the service, while the Golang backend manages the logic for shortening and redirecting URLs.

## Technologies Used

- **Frontend:**
  - [React](https://reactjs.org/)
  - [HTML5](https://developer.mozilla.org/en-US/docs/Web/Guide/HTML/HTML5)
  - [CSS3](https://developer.mozilla.org/en-US/docs/Web/CSS)
  - [JavaScript (ES6+)](https://developer.mozilla.org/en-US/docs/Web/JavaScript)

- **Backend:**
  - [Golang](https://golang.org/)
  - [MongoDB](https://www.mongodb.com/)
  - [Gorilla Mux](https://github.com/gorilla/mux)
  - [Go Modules](https://blog.golang.org/using-go-modules)

## Installation and Setup

### Prerequisites

- **Node.js** and **npm** (for the frontend)
- **Golang** (for the backend)
- **MongoDB** (for URL storage)

### Backend Setup

1. **Clone the repository and navigate to the backend directory:**

   ```bash
   git clone https://github.com/allanbrunobr/url-shortener.git
   cd url-shortener/backend

2. **Install dependencies:**
   ```bash
   go mod tidy

3. **Start the server:**
   
Make sure MongoDB is running, then execute:
```bash
   go run main.go
```
The backend will be running at http://localhost:8080.


### Configuração do Frontend

1. **Navigate to the frontend directory:**

   ```bash
   cd ../front
   ```
2. **Install dependencies:**
   ```bash
   npm install
   ```
3. **Start the development server:**
   ```bash
   npm start
   ```
   
The frontend will be accessible at http://localhost:3000.

## How to Use
1. Access the user interface at http://localhost:3000.
2. Enter the URL you want to shorten.
3. Optionally, enter a custom slug for the URL.
4. Click "Shorten URL" to generate the shortened URL.
5. Click the shortened link to be redirected to the original URL.

## Estrutura de Pastas

 ```plaintext
 url-shortener/
├── backend/                   # Backend code in Golang
│   ├── main.go                # Backend entry point
│   ├── models/                # Data models for MongoDB
│   └── ...                    # Other backend files
└── frontend/                  # Frontend code in React
    ├── src/                   # React source code
    ├── public/                # Public static files
    └── ...                    # Other frontend files
```
## Contributing
Contributions are welcome! Follow the steps below to contribute:
1. Fork the project
2. Create a branch for your feature (git checkout -b feature/new-feature)
3. Commit your changes (git commit -m 'Add new feature')
4. Push the code to your branch (git push origin feature/new-feature)
5. Open a Pull Request

## License
Distributed under the MIT License. See LICENSE for more information.

## Contact
Allan Bruno - allanbruno@gmail.com

Project Link: https://github.com/allanbrunobr/url-shortener

