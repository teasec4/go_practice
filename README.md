# Go Bank Server

A simple banking application written in Go with both CLI and HTTP server modes.

## Project Structure

```
.
├── cmd/
│   └── server/           # Server executable entry point
│       └── main.go
├── internal/
│   ├── account/          # Account domain logic
│   │   └── account.go
│   ├── bank/             # Bank management
│   │   └── bank.go
│   └── handler/          # HTTP request handlers
│       └── handler.go
├── Makefile              # Build commands
├── go.mod
└── README.md
```

## Features

- Create accounts with initial balance
- Deposit money
- Withdraw money
- Check balance
- Thread-safe concurrent operations
- HTTP API endpoints

## Installation

```bash
go mod download
```

## Usage

### Run Server

```bash
go run ./cmd/server -port 8080
```

Or using make:
```bash
make run
```

### Build

```bash
make build
```

Binary will be created at `bin/server`

## API Endpoints

### Deposit
```bash
POST /deposit
Content-Type: application/json

{
  "account_id": 1,
  "amount": 500
}
```

### Withdraw
```bash
POST /withdraw
Content-Type: application/json

{
  "account_id": 1,
  "amount": 200
}
```

### Check Balance
```bash
GET /balance?account_id=1
```

## Testing

```bash
make test
```

## Development

Format code:
```bash
make fmt
```

Run linter:
```bash
make lint
```

Clean build artifacts:
```bash
make clean
```
