# Currency Conversion

## Currency Conversion Service using Golang

### This API service was created for currency conversion with Yahoo Finance

**Used libraries:**

- [gin](https://github.com/gin-gonic)
- [gin-swagger](https://github.com/swaggo/gin-swagger)
- [gorm](https://gorm.io/docs/)
- [jwt-go](https://pkg.go.dev/gopkg.in/dgrijalva/jwt-go.v3?tab=doc)
- [godotenv](https://pkg.go.dev/github.com/joho/godotenv?tab=doc)
- [testify](https://github.com/stretchr/testify)
- [go-redis](github.com/go-redis/redis/v9)
- [pq](github.com/lib/pq)
- [crypto](golang.org/x/crypto)
- [swag](github.com/swaggo/swag)
- [swaggo-files](github.com/swaggo/files)
- [gin-swagger](github.com/swaggo/gin-swagger)
- [go-sqlmock](github.com/DATA-DOG/go-sqlmock)

### Run locally

```sh
docker build .           # docker build
docker-compose up        # docker-compose up (Run postgres / redis and other volumes)
docker-compose down      # docker-compose down (Shutdown postgres / redis and other volumes)
```

### Documentation

```sh
SWAGGER DOCUMENTATION: http://localhost:8080/swagger/index.html
```

#### If you want to change the fee percentage, you can change the `FEE_PERCENTAGE` value in the .env
