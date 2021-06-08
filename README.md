# Technologies

- Golang;
- Docker;
- PostgreSQL

# Libraries

- [PGX](https://github.com/swaggo/swag)
- [Zerolog](https://github.com/rs/zerolog)
- [PGX](https://github.com/jackc/pgx)
- [Mux](https://github.com/gorilla/mux)
- [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [JWT-Go](https://github.com/dgrijalva/jwt-go)
- [Google uuid](https://github.com/google/uuid)
- [Golang-migrate](https://github.com/golang-migrate/migrate)
- [Moq](https://github.com/matryer/moq)

# Getting Started

### Starting the application
```zsh
  make dev-up
```

### Shutting application
```zsh
   make dev-down
```

The server will be listening on localhost:8080.
# End Points

### The complete API documentation is available at /swagger.

Demo: http://localhost:8080/docs/v1/swagger/index.html

### `/api/v1/accounts - POST`

```json
{
  "name": "gabriel gabriel",
  "secret": "45478494894",
  "cpf": "57364376003
}
```

### Details

- `name` field must have at least 8 characters
- `secret` field must have at least 8 characters
- `cpf` field must be a valid cpf

### `/api/v1/accounts - Get`

### Response

```json
[
  {
    "id": "e91e8ac1-94f2-4dff-bcbf-e5941fc253ad",
    "name": "hu3hu3hu3dsadasd",
    "cpf": "02732567000",
    "balance": 0,
    "created_at": "2021-06-05T04:03:32.984852Z"
  }
]
```

### `/api/v1/accounts/{account_id}/balance - Get`

### Response

```json
{
  "id": "abce8b02-5f3a-4f2c-96a7-964e37d0dc08",
  "balance": 0
}
```

### `/api/v1/login - POST`

```json
{
  "secret": "45478494894",
  "cpf": "57364376003"
}
```

### Response

```json
{
  "token": "jwt token"
}
```

### `/api/v1/transfers - POST`

```json
{
  "account_destination_id": "abce8b02-5f3a-4f2c-96a7-964e37d0dc08",
  "amount": 100,
}
```

### `/api/v1/transfers - Get` (Bearer token required)

### Response

```json
[
  {
     "id": "e91e8ac1-94f2-4dff-bcbf-e5941fc253ad",
     "account_destination_id": "abce8b02-5f3a-4f2c-96a7-964e37d0dc08",
     "amount": 100,
     "created_at": "2021-06-04T23:22:00.622329Z"
  }
]
```

## Testing

### Run tests
```zsh
  make test
```

