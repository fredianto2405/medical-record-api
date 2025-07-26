# Electronic Medical Record API
This project is a comprehensive clinic or hospital management system API built with Go. 
It is designed to handle various core operations such as patient registration, medical records, billing, and treatment workflows.

## Features
- User Authentication with JWT
- Clinic Management
- Doctor Management
- Nurse Management
- Insurance Management
- Patient Management
- Medicine Management
- Treatment Management
- Payment Management
- Medical Record Management
- User Management

## Tech Stack
- Golang
- Gin Framework
- PostgreSQL
- JWT for Authentication

## Get Started

### Prerequisites
- Go 1.24.4
- PostgreSQL
- Git
- Insomnia 11.3.0

### Clone & Run

```bash
git clone https://github.com/fredianto2405/medical-record-api.git
cd medical-record-api
go mod tidy
go run cmd/main.go
```

### Environment Variables
Create a `.env` file and fill with the following:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=yourpwd
DB_NAME=yourdb
PORT=8080
JWT_SECRET=your_jwt_secret
```

## API Documentation
Check on directory docs/electronic_medical_record.yaml, then import to Insomnia