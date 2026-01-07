# Vyoma – Compliance Management Backend

Vyoma is a **compliance management backend system** built in Go that automatically assigns, tracks, and manages statutory compliances for businesses based on **state, industry, and turnover**.

This project is designed as a **real-world, production-grade backend**, following clean architecture principles and transactional safety.

---

## 🚀 Project Vision

Businesses often struggle to track regulatory compliances that vary by:
- State
- Industry
- Turnover size

Vyoma solves this by acting as a **compliance engine**:
- Businesses are registered once
- Applicable compliance rules are auto-assigned
- Status, due dates, and authorities are tracked
- The system remains auditable, extensible, and reliable

---

## 🧱 Architecture Overview

The backend follows a **clean layered architecture**:

HTTP (Handlers)
↓
Services (Business Logic & Transactions)
↓
Repositories (Database Queries)
↓
PostgreSQL (Supabase)


Each layer has a **single responsibility**, making the system:
- Easy to maintain
- Easy to extend
- Safe to refactor
- Testable in isolation

---

## 🛠 Tech Stack

- **Language**: Go
- **Router**: Chi
- **Database**: PostgreSQL (Supabase)
- **DB Driver**: pgx v5
- **Migrations**: golang-migrate
- **Containerization**: Docker & Docker Compose
- **API Style**: REST
- **Environment Management**: `.env` + Docker

---

## 📂 Project Structure

cmd/
└── server/
└── main.go

internal/
├── db/
├── models/
├── repositories/
├── services/
├── handlers/
├── routes/

migrations/
Dockerfile
docker-compose.yml
go.mod
go.sum



---

## 🗄 Database Schema

### businesses
Stores registered businesses.

**Fields**
- `id` (UUID)
- `name`
- `state`
- `industry_type`
- `turnover_range`
- `created_at`

---

### compliance_rules
Master table defining all compliance rules.

**Fields**
- `id`
- `state`
- `industry_type`
- `turnover_min`
- `turnover_max`
- `rule_name`
- `frequency` (monthly / yearly)
- `authority`
- `penalty_amount`

---

### business_compliance_map
Maps businesses to applicable compliance rules.

**Design**
- Composite Primary Key: `(business_id, rule_id)`
- Surrogate `id` column for API ergonomics

**Fields**
- `id` (UUID, unique)
- `business_id`
- `rule_id`
- `status` (pending / completed / overdue)
- `due_date`

---

## ✅ Features Implemented

- Create Business API
- Get Business by ID
- Automatic compliance rule assignment
- Transactional safety (rollback on failure)
- List compliances for a business
- Dockerized runtime environment

---

## 🔄 Transaction Flow (Important)

When a business is created:

1. Start DB transaction
2. Insert business
3. Fetch applicable compliance rules
4. Assign rules to business
5. Commit transaction

If **any step fails**, the entire operation is rolled back.

This guarantees **data consistency** at all times.

---

## 🌐 API Endpoints

### Create Business
POST /businesses

### Get Business
GET /businesses/{id}

### List Business Compliances
GET /businesses/{id}/compliances

---

## 🐳 Running with Docker

### Build & Run
```bash
docker compose up --build

The server runs on:
http://localhost:8080

🧪 Local Development (without Docker)

go run cmd/server/main.go

Make sure environment variables are set.

🔐 Environment Variables

DATABASE_URL=postgresql://...
⚠️ .env is intentionally not committed.

📈 Roadmap (Next Features)

Mark compliance as completed

Overdue compliance detection (scheduler)

Notification system

Authentication & multi-tenancy

Role-based access control

CI/CD pipeline

Production deployment

🎓 What This Project Teaches

By building this system, you gain hands-on experience with:

Clean architecture in Go

Transactional database design

Schema evolution strategies

Docker-based workflows

Real-world backend patterns

Designing domain-driven systems

📜 License

MIT License

👤 Author

Built with care as a serious backend engineering project.


---
