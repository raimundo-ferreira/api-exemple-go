# 📖 Exemple: API with Golang

Example API with Golang using Gin, PostgreSQL, PGX, JWT and Swagger.

## Packages
```bash
go get -u github.com/gin-gonic/gin
go get -u github.com/jackc/pgx/v5
go get -u github.com/golang-jwt/jwt/v5

go get github.com/joho/godotenv/cmd/godotenv

go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

## Quick start

1. Create database in postgresql and execute the scrip "internal/database/schema.sql"
2. Open the `.env` and fill it with your environment values.
3. Run the project with the commands:

```bash
swag init
go run .

# Process:
#   - Generate API docs by Swagger
#   - Execute api
```

4. Go to your API Docs page: [127.0.0.1:5000/swagger/index.html](http://127.0.0.1:5000/swagger/index.html)
5. Use access key "42975313-085b-4c09-90c0-11fe6ba6de5d" to register a user
