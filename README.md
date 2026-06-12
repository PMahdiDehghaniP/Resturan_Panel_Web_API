# Restaurant Panel Web API

A Go REST API for a restaurant management panel. The service currently exposes
customer and food endpoints over Gin, stores data in PostgreSQL, and writes
structured logs through a configurable Zap or Zerolog logger.

The database schema is broader than the current HTTP API. It already models
customers, addresses, tables, discounts, invoices, orders, payments, employees,
food categories, and food items.

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Runtime Flow](#runtime-flow)
- [Configuration](#configuration)
- [Local Setup](#local-setup)
- [Database Setup](#database-setup)
- [Run the API](#run-the-api)
- [API Overview](#api-overview)
- [Logging](#logging)
- [Development Notes](#development-notes)
- [Troubleshooting](#troubleshooting)
- [More Documentation](#more-documentation)

## Features

- Versioned HTTP API under `/api/v1`.
- PostgreSQL connection pool configured from YAML.
- Customer listing with pagination.
- Food listing with optional category filtering.
- Food creation with JSON request validation.
- SQL schema and seed data for restaurant panel domain tables.
- Structured file logging with Zap or Zerolog.
- Docker Compose infrastructure for PostgreSQL, pgAdmin, Elasticsearch, Kibana,
  and Filebeat.

## Tech Stack

- Language: Go 1.25.7, based on `go.mod`
- HTTP framework: Gin
- Database: PostgreSQL
- SQL driver: `github.com/lib/pq`
- Configuration: Viper plus `godotenv`
- Logging: Zap or Zerolog
- Log rotation: Lumberjack for Zap logger
- Infrastructure: Docker Compose

## Project Structure

```text
.
|-- main.go
|-- api/
|   |-- api.go
|   |-- routes/
|   `-- handlers/
|-- config/
|   |-- app.env
|   |-- config_dev.yml
|   |-- config_docker.yml
|   `-- config_production.yml
|-- data/
|   |-- db/
|   `-- sql-scripts/
|-- docker/
|   `-- docker-compose.yaml
|-- logger/
|-- models/
|-- ResturnDB_Schema.sql
`-- mockdata_migration.sql
```

Important directories:

- `api/`: Gin server setup, route registration, and HTTP handlers.
- `config/`: environment selector and YAML configuration files.
- `data/db/`: PostgreSQL initialization and connection pool lifecycle.
- `data/sql-scripts/`: SQL strings used by handlers.
- `logger/`: logger interface plus Zap and Zerolog implementations.
- `models/`: request, response, and shared DTO structs.
- `docker/`: local infrastructure services.

## Runtime Flow

`main.go` is intentionally small:

1. Loads application configuration with `config.GetConfig()`.
2. Initializes PostgreSQL with `db.InitPostgresDB(appConfig)`.
3. Defers database shutdown with `db.ClosePostgresDB()`.
4. Starts the Gin HTTP server with `api.InitApiServer()`.

The API server:

1. Reads the configured server port.
2. Creates a Gin engine with request logging and recovery middleware.
3. Registers routes under `/api/v1`.
4. Starts listening on the configured port.

## Configuration

The application always loads `./config/app.env` first.

```env
APP_TYPE=dev
CONFIG_TYPE=yml
```

`APP_TYPE` selects which YAML file is loaded:

| APP_TYPE | Config file |
| --- | --- |
| `dev` | `config/config_dev.yml` |
| `docker` | `config/config_docker.yml` |
| `production` | `config/config_production.yml` |

The development configuration currently uses:

```yaml
server:
  Port: 8080
  RunMode: debug
postgres:
  host: localhost
  port: 5433
  user: postgres
  password: 12345678
  dbName: restaurant_db
  sslMode: disable
logger:
  filePath: logs/logs.log
  encoding: json
  level: debug
  loggerName: Zap
```

The active logger is selected with `logger.loggerName`:

- `Zap`
- `Zero`

## Local Setup

Requirements:

- Go compatible with the version in `go.mod`
- Docker and Docker Compose
- PostgreSQL client tools such as `psql`

Install dependencies:

```bash
go mod download
```

Start infrastructure:

```bash
docker compose -f docker/docker-compose.yaml up -d postgres pgadmin
```

Optional logging infrastructure:

```bash
docker compose -f docker/docker-compose.yaml up -d elasticsearch kibana filebeat
```

Service ports from the current compose file:

| Service | Host port | Container port |
| --- | ---: | ---: |
| PostgreSQL | `5433` | `5432` |
| pgAdmin | `8000` | `80` |
| Elasticsearch | `9200` | `9200` |
| Kibana | `5601` | `5601` |

## Database Setup

The compose file creates the database configured by `config/config_dev.yml`:

- Host: `localhost`
- Port: `5433`
- User: `postgres`
- Password: `12345678`
- Database: `restaurant_db`

Apply schema:

```bash
psql "postgres://postgres:12345678@localhost:5433/restaurant_db?sslmode=disable" \
  -f ResturnDB_Schema.sql
```

Apply seed data:

```bash
psql "postgres://postgres:12345678@localhost:5433/restaurant_db?sslmode=disable" \
  -f mockdata_migration.sql
```

The seed data includes food categories, food items, customers, employees, tables,
orders, invoices, discounts, and payment transactions.

## Run the API

Run locally:

```bash
go run .
```

Default development URL:

```text
http://localhost:8080
```

Health-check style request using an implemented endpoint:

```bash
curl "http://localhost:8080/api/v1/foods/getfoods"
```

## API Overview

Base path:

```text
/api/v1
```

Implemented endpoints:

| Method | Path | Description |
| --- | --- | --- |
| `GET` | `/customers/getall` | List customers with pagination. |
| `GET` | `/foods/getfoods` | List foods, optionally filtered by category name. |
| `POST` | `/foods/createfood` | Create a food item. |

Example: list customers

```bash
curl "http://localhost:8080/api/v1/customers/getall?page=1&pageSize=10"
```

Example: list foods by category

```bash
curl "http://localhost:8080/api/v1/foods/getfoods?foodCategory=Pizza"
```

Example: create a food item

```bash
curl -X POST "http://localhost:8080/api/v1/foods/createfood" \
  -H "Content-Type: application/json" \
  -d '{
    "foodName": "Mushroom Burger",
    "description": "Beef burger with mushrooms and cheese",
    "calories": 780,
    "imagePath": "/images/mushroom-burger.jpg",
    "unitPrice": 210000,
    "inventory": 25,
    "foodCategory": 2
  }'
```

See [API documentation](docs/API.md) for response shapes and parameter details.

## Logging

Logs are written to the file configured by `logger.filePath`, currently:

```text
logs/logs.log
```

Zap logging uses JSON encoding and Lumberjack rotation:

- Maximum file size: 10 MB
- Maximum age: 5 days
- Maximum backups: 10
- Local timestamps enabled
- Compression enabled

Filebeat can ship `logs/` into Elasticsearch when the optional logging services
are running.

## Development Notes

- Route registration starts in `api/routes/routes.go`.
- Handler logic is split by domain under `api/handlers/`.
- SQL strings are stored under `data/sql-scripts/`.
- PostgreSQL connection pooling is configured in `data/db/postgres.go`.
- Current API responses use a mix of explicit JSON tags and default Go field
  names. See `models/` before changing response contracts.
- `config/config_docker.yml` and `config/config_production.yml` currently contain
  older database names/passwords that do not fully match the development compose
  file. Review them before using those modes.

## Troubleshooting

### `APP_TYPE or CONFIG_TYPE environment variable not set`

The app could not load `config/app.env`. Run commands from the repository root
or make sure `./config/app.env` exists.

### `config file not found`

`APP_TYPE` does not map to a known config file, or `CONFIG_TYPE` does not match
the YAML extension. For local development, use:

```env
APP_TYPE=dev
CONFIG_TYPE=yml
```

### `connection refused` when starting the API

PostgreSQL is not running or the port does not match the active config. For the
development config, PostgreSQL must be reachable on `localhost:5433`.

### `pq: relation "customer" does not exist`

The database schema has not been applied. Run `ResturnDB_Schema.sql`, then load
`mockdata_migration.sql` if you want sample data.

### Empty food category filter result

`foodCategory` filters by category name, not category id. Example:

```text
foodCategory=Pizza
```

## More Documentation

- [API documentation](docs/API.md)
- [Database documentation](docs/DATABASE.md)
