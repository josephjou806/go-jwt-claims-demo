Go JWT Claims Demo

A minimal Go microservice using Gin and JWT authentication, plus a simple static web client (served via GitHub Pages).

🌐 Live web client: https://josephjou806.github.io/go-jwt-claims-demo/

Features

REST API with Gin

JWT-based authentication middleware

/claims/{id} endpoint (mocked in-memory repository)

Configurable via environment variables

Dependency injection–friendly code structure

Unit tests for JWT and claim lookups

Static web client hosted on GitHub Pages

Project Layout
.
├── cmd/server/main.go          # API entrypoint
├── internal/                   # business logic and layers
│   ├── config/                 # env config loader
│   ├── handlers/               # HTTP handlers
│   ├── middleware/             # JWT + CORS
│   ├── models/                 # domain models
│   ├── repository/             # in-memory claim repo
│   ├── server/                 # router wiring
│   ├── services/               # business services
│   └── token/                  # JWT manager
├── tests/                      # Go unit tests
├── docs/                       # static web client for GitHub Pages
└── go.mod / go.sum             # module dependencies

Running Locally
1. Clone & install deps
git clone https://github.com/josephjou806/go-jwt-claims-demo.git
cd go-jwt-claims-demo
go mod tidy

2. Run the API
export PORT=8080
export JWT_SECRET=dev-secret-change-me
export CORS_ALLOWED_ORIGINS=http://localhost:5500
go run ./cmd/server

3. Run the web client
python3 -m http.server 5500


Then visit:
👉 http://localhost:5500/docs/index.html

GitHub Pages Deployment

The docs/ folder contains the web client. GitHub Pages is enabled for main /docs.

Visit: https://josephjou806.github.io/go-jwt-claims-demo/

The client reads API base URL from docs/config.json

{ "API_BASE": "https://<your-api-host>" }


⚠️ The API must be served over HTTPS and configured with:

export CORS_ALLOWED_ORIGINS=https://josephjou806.github.io

Running Tests
go test ./... -v

Example API Usage
Login
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"demo","password":"demo"}'


Response:

{ "token": "eyJhbGciOi..." }

Get Claim
curl -H "Authorization: Bearer <TOKEN>" \
  http://localhost:8080/claims/1001


Response:

{
  "id": "1001",
  "memberId": "M-001",
  "ndc": "0002-8215-01",
  "amount": 15.75,
  "status": "PAID",
  "fillDate": "2025-01-15T12:00:00Z"
}

License

MIT
