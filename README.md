# Simple Bank Backend

A production-ready banking backend built with **Go**, **gRPC**, **JWT**, **Redis**, and **PostgreSQL**. 
Deployed on **AWS EKS** with a full CI/CD pipeline.

---

## ✨ Features

- 🔐 **Authentication & Authorization**: JWT with refresh tokens and session management.
- 👤 **User Management**: Registration, login, and profile updates.
- 🏦 **Banking Operations**: Create accounts, record transfers, and view transaction history.
- ⚙️ **Background Workers**: Asynchronous task processing with Redis and email workers.
- 🗄️ **Database**: PostgreSQL with SQLC for type-safe queries and migrations.
- 📡 **API**: gRPC and REST (via gRPC-Gateway) with Swagger documentation.
- 🐳 **Containerisation**: Docker and Docker Compose for local development.
- ☁️ **Cloud Deployment**: Kubernetes manifests for AWS EKS.
- 🔄 **CI/CD**: Automated builds and deployments with GitHub Actions.

---

## 🛠️ Tech Stack

| Component | Technology |
|-----------|------------|
| Language | Go 1.22+ |
| API | gRPC + gRPC-Gateway (REST) |
| Database | PostgreSQL 15 |
| SQL | SQLC (type-safe queries) |
| Migrations | golang-migrate/migrate |
| Cache / Worker | Redis |
| Authentication | JWT (with refresh tokens) |
| Containerisation | Docker, Docker Compose |
| Orchestration | Kubernetes (AWS EKS) |
| CI/CD | GitHub Actions |
| Documentation | Swagger (OpenAPI) |

---

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/eugenius-watchman/golang_simplebank.git
cd golang_simplebank

2. Run with Docker Compose (recommended for local development)
docker-compose up
This starts:

PostgreSQL on localhost:5432

Redis on localhost:6379

The gRPC server on localhost:9090

The REST gateway on localhost:8080

Swagger UI on localhost:8080/swagger/

3. Run locally (without Docker)
# Install dependencies
go mod download

# Set up environment variables
cp app.env.example app.env

# Edit app.env with your database and Redis credentials
# Run database migrations
make migrateup

# Start the server
go run main.go
📡 API Endpoints
gRPC
Method	Description
CreateUser	Register a new user
LoginUser	Authenticate and get JWT tokens
UpdateUser	Update user profile (with validation)
GetUser	Get user details
CreateAccount	Open a new bank account
GetAccount	Get account details
ListAccounts	List all accounts for a user
CreateTransfer	Transfer money between accounts
REST (via gRPC-Gateway)
Endpoint	Method	Description
/v1/users	POST	Create a new user
/v1/users/{id}	GET	Get user details
/v1/users/{id}	PATCH	Update user
/v1/login	POST	User login
/v1/accounts	POST	Create account
/v1/accounts/{id}	GET	Get account
/v1/transfers	POST	Transfer funds
📚 API Documentation
Swagger UI is available at:

http://localhost:8080/swagger/
🧪 Running Tests
make test
📂 Project Structure
.
├── api/           # gRPC and REST API handlers
├── db/            # Database migrations and queries
│   ├── migration/ # SQL migration files
│   ├── query/     # SQLC queries
│   └── sqlc/      # Generated type-safe Go code
├── doc/           # Swagger documentation
├── eks/           # Kubernetes manifests for AWS EKS
├── gapi/          # gRPC server implementation
├── pb/            # Protocol Buffer definitions and generated code
├── proto/         # .proto files
├── token/         # JWT token management
├── util/          # Utility functions
├── val/           # Request validators
├── worker/        # Background task workers (Redis)
├── .github/workflows/ # CI/CD pipelines
├── Dockerfile
├── docker-compose.yaml
├── Makefile
└── README.md
☁️ Deployment
Docker
Build and run the Docker image:

docker build -t simplebank .
docker run -p 8080:8080 -p 9090:9090 simplebank
Kubernetes (AWS EKS)
Apply the Kubernetes manifests:

kubectl apply -f eks/
This deploys the application with:

2 replicas for high availability

LoadBalancer service

PostgreSQL and Redis dependencies (or use AWS RDS and ElastiCache)

🔄 CI/CD Pipeline
This project uses GitHub Actions for continuous integration and deployment:

Workflow	Trigger	Actions
CI	Push to main	Run tests, build Docker image
CD	Release creation	Build and push to AWS ECR, deploy to EKS
📌 Future Improvements
Add observability (Prometheus, Grafana)

Implement rate limiting

Add integration tests with testcontainers

Support for multi-currency accounts

👤 Author
Eugene Darrah-Gblorkpor – GitHub







