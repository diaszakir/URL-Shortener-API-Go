# URL Shortener API Project

One of the project of golang [list](https://github.com/diaszakir/Golang-Projects/)

## Description

URL Shortener service - simple service for shortening URL

Includes redirecting to website

## How to launch
Copy repository to your local machine using command:

```
git clone https://github.com/diaszakir/URL-Shortener-API-Go
```

In case if you do not have go.mod and go.sum

```
go get -u github.com/gin-gonic/gin
go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
go get -u github.com/diaszakir/base62-go
```

After

```
go run cmd/app/main.go
```

Loading documentation in case if you changed structure

```
swag init -g cmd/app/main.go
```

To access swagger go to `localhost:8080/swagger/index.html`

## API Endpoints
- `GET /health` - checks API
- `GET /:code` - Redirecting from URL
- `GET /info/:code` - Get info about shortered URL
- `POST /shorten` - Short given URL