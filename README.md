#  gRPC Report Generation Service

It is a lightweight, production-style Go service built for Appversal's machine task. It exposes a gRPC API for report generation, includes a scheduled cron job, and persists reports using PostgreSQL.

---

##  Features

-  gRPC endpoint: `GenerateReport(UserID)` → `(ReportID, Error)`
-  Cron job: runs every 10 seconds to generate reports for predefined users
-  Health check endpoint: `HealthCheck()` → `"SERVING"`
-  Zap-structured logging with timestamps
-  PostgreSQL storage using UUIDs
-  Clean architecture & modular layout

---

##  Project Structure
.
├── main.go
├── .env
├── config/
│ └── config.go
├── internal/
│ ├── model/
│ │ └── report.go
│ ├── repository/
│ │ └── repository.go
│ └── service/
│ └── server.go
├── pkg/
│ └── scheduler.go
├── proto/
│ ├── report.proto
│ └── report.pb.go (and other generated files)
├── README.md
└── SCALE_DESIGN.md


---

##  Setup Instructions

### 1. Clone the Repo
```bash
git clone https://github.com/Prototype-1/grpc-report-service.git
cd grpc-report-service

```

### .env

GRPC_PORT=50050
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=YOUR_PASSWORD
DB_NAME=grpc_reports_db

##  Run the Service

go run main.go

### Testing gRPC with Postman

  1.  Open Postman → New > gRPC Request

  2.  Connect to: localhost:50050

  3. Load the .proto file (proto/report.proto)

a. GenerateReport Method
{
  "user_id": "aswin"
}

b. HealthCheck Method
{}

### Cron Job Behavior

    Every 10 seconds, the cron scheduler invokes GenerateReport for:

        user1, user2, user3

    Reports are persisted in PostgreSQL with timestamps and UUIDs.

## Dependencies

    Go 1.24+

    gRPC (google.golang.org/grpc)

    GORM + PostgreSQL

    Zap logging (go.uber.org/zap)

    Cron scheduler (github.com/robfig/cron/v3)

    UUIDs (github.com/google/uuid)

    Env loader (github.com/joho/godotenv)

### Bonus

See SCALE_DESIGN.md for a detailed system design plan to scale this service for 10,000 concurrent gRPC requests per second across multiple data centers.

## Author

Aswin P Raghu
Machine Task Submission for Appversal