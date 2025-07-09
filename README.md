# Web Shop — Online Store in Go
A simple online store written in Go using the Gin framework and HTMX for dynamic HTML updates without JavaScript.

## Overview

This is a learning project for a web shop built with:
- User authentication via JWT.
- Support for multiple databases.
- Product management, cart, orders, categories.
- WebSocket-based interaction.
- HTMX-powered dynamic UI rendering.

The code follows the **Clean Architecture** approach and separates concerns into layers: `handlers`, `services`, `storage`. It uses the **Repository design pattern** to abstract data access logic and decouple business logic from the database implementation.

## Technologies Used

- **Go 1.20+**
- **Gin Gonic**
- **HTMX**
- **JWT**
- **WebSocket**
- **PostgreSQL / MSSQL**

## Repository Design Pattern

This project uses the **Repository design pattern**, which provides a clean separation between domain logic and data access logic.

### Key Benefits:
- Decouples business logic from the database.
- Makes it easy to switch between different data sources (e.g., PostgreSQL ↔ MSSQL).
- Simplifies testing by allowing mock implementations.

Each service depends on an interface that defines how it interacts with the data layer. The actual implementation is injected at runtime based on the configured database driver.

Example:
```go
type ReadWriter interface {
    AddUser(ctx context.Context, user *User) error
    GetUser(ctx context.Context, email string) (*User, error)
}
```

## How to Run

### 1. Install dependencies

```bash
git clone https://github.com/artyomkorchagin/web-shop-laba.git
cd web-shop-laba
go mod tidy
```

### 2. Create a config file

Place a config like `config/psql.yaml` in the `config` folder.

Example:

```yaml
dbms: psql
driver: pgx
server:
  host: localhost
  port: 5432
database:
  name: shop_db
  user: admin
  password: pass
  sslmode: disable
logger:
  level: info
  format: json
```

### 3. Run the app

```bash
go run cmd/main.go -dbms psql
```

or

```bash
go run cmd/main.go -dbms mssql
```

The app will be available at: `http://localhost:3000`
