# Vyoma API 🚀

Production-ready backend API built with Go.

## 🧱 Tech Stack

- Go (Chi)
- PostgreSQL (Supabase)
- Docker & Docker Compose
- pgx (PostgreSQL driver)

---

## 📁 Project Structure

.
├── cmd/server # App entry point
├── internal # Business logic
├── db # Database connection
├── Dockerfile
├── docker-compose.yml
├── .env.example

yaml
Copy code

---

## ⚙️ Environment Variables

Create a `.env` file in the root:

DATABASE_URL=your_database_url_here
PORT=8080

yaml
Copy code

> ⚠️ Never commit `.env` files

---

## 🐳 Run with Docker

### Build image
```bash
docker build -t vyoma-api .
Run using docker-compose
bash
Copy code
docker compose up
API will be available at:

arduino
Copy code
http://localhost:8080
🩺 Health Check
bash
Copy code
GET /health
Returns:

json
Copy code
{ "status": "ok" }
📌 Notes
Uses multi-stage Docker builds

Distroless runtime image

Externalized configuration

Ready for production deployment

yaml
Copy code

---

## 🧠 Why This README Matters

This tells reviewers:
- You understand infra
- You understand security
- You think like a backend engineer
- You’re not a tutorial copier

This alone separates you from **90% of repos**.

---

## ✅ Commit Everything

```bash
git add docker-compose.yml README.md
git commit -m "Add docker-compose and project README"
git push