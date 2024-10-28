# URL Shortener App

This is a URL shortener application that stores URLs in a SQLite3 database and returns a shortened URL for the user to use. The backend server is built using Golang.

## Prerequisites

- Go (version 1.16 or higher)
- SQLite3

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/urlshortener.git
    cd urlshortener
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

## Database Setup

1. Create the SQLite3 database:
    ```sh
    sqlite3 urls.db < schema.sql
    ```

2. The `schema.sql` file should contain the following SQL to create the necessary table:
    ```sql
    CREATE TABLE urls (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        original_url TEXT NOT NULL,
        short_url TEXT NOT NULL UNIQUE
    );
    ```

## Running the App

1. Run the application:
    ```sh
    go run main.go
    ```

2. The server will start on `http://localhost:8080`.

## Usage

### Shorten a URL

To shorten a URL, send a POST request to `http://localhost:8080/shorten` with the original URL in the request body.

Example using `curl`:
```sh
curl -X POST -d '{"url": "https://www.example.com"}' -H "Content-Type: application/json" http://localhost:8080/shorten
```

### Redirect to Original URL

To redirect to the original URL, use the shortened URL provided by the app.

Example:
```sh
http://localhost:8080/{short_url}
```

## API Endpoints

- `POST /shorten`: Accepts a JSON body with the original URL and returns a shortened URL.
- `GET /{short_url}`: Redirects to the original URL associated with the shortened URL.

## Example

1. Shorten a URL:
    ```sh
    curl -X POST -d '{"url": "https://www.example.com"}' -H "Content-Type: application/json" http://localhost:8080/shorten
    ```

    Response:
    ```json
    {
        "short_url": "http://localhost:8080/abc123"
    }
    ```

2. Use the shortened URL to redirect:
    ```sh
    http://localhost:8080/abc123
    ```

    This will redirect to `https://www.example.com`.

## License

This project is licensed under the MIT License.
