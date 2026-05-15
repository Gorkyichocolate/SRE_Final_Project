# Backend Testing

## Run PostgreSQL

```bash
sudo systemctl start postgresql
```

---

## Create Database

```bash
sudo -u postgres psql
```

```sql
CREATE DATABASE ecommerce;

ALTER USER postgres WITH PASSWORD 'postgres';

\q
```

---

## Run Migrations

```bash
sudo -u postgres psql -d ecommerce -f migrations/init.sql
```

---

## Configure Environment

Create `.env` file in project root:

```env
PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=ecommerce

JWT_SECRET=super-secret-key
```

---

## Install Dependencies

```bash
go mod tidy
```

---

## Run Backend

```bash
go run ./cmd/api
```

Backend starts on:

```text
http://localhost:8080
```

---

# API Testing

---

## Register User

```bash
curl -X POST http://localhost:8080/register \
-H "Content-Type: application/json" \
-d '{
  "username":"biba",
  "email":"biba@gmail.com",
  "password":"123456"
}'
```

Expected response:

```json
{
  "message":"user created"
}
```

---

## Login User

```bash
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{
  "email":"biba@gmail.com",
  "password":"123456"
}'
```

Expected response:

```json
{
  "token":"JWT_TOKEN"
}
```

Copy the JWT token.

---

## Get Products

```bash
curl http://localhost:8080/products
```

---

## Get Product By ID

```bash
curl http://localhost:8080/products/1
```

---

## Protected Route

Replace `JWT_TOKEN` with real token.

```bash
curl http://localhost:8080/protected \
-H "Authorization: Bearer JWT_TOKEN"
```

Expected response:

```json
{
  "message":"protected route"
}
```

---

## Create Order

```bash
curl -X POST http://localhost:8080/orders \
-H "Authorization: Bearer JWT_TOKEN" \
-H "Content-Type: application/json" \
-d '{
  "product_id":1,
  "quantity":2
}'
```

Expected response:

```json
{
  "message":"order created"
}
```

---

## Get Orders

```bash
curl http://localhost:8080/orders \
-H "Authorization: Bearer JWT_TOKEN"
```

---

## Health Check

```bash
curl http://localhost:8080/health
```

Expected response:

```json
{
  "status":"ok"
}
```

---

## Metrics Endpoint

```bash
curl http://localhost:8080/metrics
```

Expected metrics:

```text
http_requests_total
http_request_duration_seconds
orders_created_total
auth_requests_total
heavy_requests_total
```

---

## Heavy Endpoint

```bash
curl http://localhost:8080/heavy
```

Expected response:

```json
{
  "message":"heavy request completed"
}
```
