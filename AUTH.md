# Authentication & Authorization Guide

## Overview

This API now includes complete authentication and authorization system with JWT tokens, password hashing, and role-based access control.

## Features

- **User Signup**: Register new users with email and password
- **User Login**: Authenticate and receive JWT token
- **Password Hashing**: Bcrypt for secure password storage
- **JWT Tokens**: Stateless authentication with TokenExpiryTime configurability
- **Request Tracking**: Every request gets a unique ID for tracing
- **Role-Based Access Control**: Admin and User roles with middleware enforcement
- **Protected Routes**: Secure endpoints that require authentication
- **Error Standardization**: Consistent error format across all endpoints

## New Endpoints

### Public (No Authentication Required)

#### Sign Up
```
POST /auth/signup
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "SecurePass123",
  "dob": "1990-05-15"
}

Response: 201 Created
{
  "message": "User registered successfully",
  "email": "john@example.com"
}
```

#### Login
```
POST /auth/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "SecurePass123"
}

Response: 200 OK
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "role": "user",
  "token": "eyJhbGciOiJIUzI1NiIs..."
}
```

### Protected (Authentication Required)

#### Get Current User Profile
```
GET /user/profile
Authorization: Bearer <token>

Response: 200 OK
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "role": "user",
  "dob": "1990-05-15"
}
```

### Admin Only (Requires Admin Role)

```
GET /admin/* (add your admin routes here)
Authorization: Bearer <token>
```

## Database Schema

The `users` table now has:
- `id` (INT, Auto-increment, Primary Key)
- `name` (VARCHAR(100))
- `email` (VARCHAR(255), Unique)
- `password_hash` (VARCHAR(255))
- `role` (ENUM: 'user', 'admin')
- `dob` (DATE)
- `created_at` (TIMESTAMP)
- `updated_at` (TIMESTAMP)

## Middleware

### RequestIDMiddleware
- Generates unique UUID for each request
- Stores in `X-Request-ID` header
- Available in context as `request_id`

### AuthMiddleware
- Validates JWT token from Bearer header or cookie
- Extracts user claims
- Stores user info in context (user_id, user_email, user_role)
- Returns 401 if token is missing/invalid

### RoleMiddleware
- Checks if user has required role
- Returns 403 if unauthorized
- AdminOnly() helper for admin-only routes

## Configuration

### JWT Secret (config/jwt.go)
```go
config.JWT.Secret = "your-secret-key"
config.JWT.ExpiryHours = 24
```

**IMPORTANT**: Change the secret in production!

### Password Hashing
Uses bcrypt with default cost (10 iterations)

## Error Response Format

All errors return consistent format:
```json
{
  "error": {
    "message": "Error description",
    "code": "ERROR_CODE",
    "request_id": "<uuid>"
  }
}
```

## HTTP Status Codes

- `200` - Success
- `201` - Created
- `400` - Bad Request (validation error)
- `401` - Unauthorized (missing/invalid token or credentials)
- `403` - Forbidden (insufficient permissions)
- `404` - Not Found
- `500` - Internal Server Error

## Example Workflow

1. **Sign Up**
   ```bash
   curl -X POST http://localhost:3000/auth/signup \
     -H "Content-Type: application/json" \
     -d '{
       "name":"Alice",
       "email":"alice@example.com",
       "password":"Pass123!",
       "dob":"1995-03-20"
     }'
   ```

2. **Login**
   ```bash
   curl -X POST http://localhost:3000/auth/login \
     -H "Content-Type: application/json" \
     -d '{
       "email":"alice@example.com",
       "password":"Pass123!"
     }'
   ```
   Response includes `token`

3. **Access Protected Route**
   ```bash
   curl -X GET http://localhost:3000/user/profile \
     -H "Authorization: Bearer <TOKEN_FROM_LOGIN>"
   ```

## Security Best Practices

✅ **Implemented**
- Bcrypt password hashing
- JWT for stateless auth
- HttpOnly secure cookies (optional)
- Request ID tracking
- Role-based access control
- Structured error messages

⚠️ **Production Checklist**
- [ ] Change JWT secret in environment variable
- [ ] Enable HTTPS (set cookie Secure=true)
- [ ] Implement rate limiting
- [ ] Add CORS configuration
- [ ] Use environment variables for secrets
- [ ] Implement logout/token blacklist
- [ ] Add password reset flow
- [ ] Implement 2FA for admins
