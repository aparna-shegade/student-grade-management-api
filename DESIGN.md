# System Design Document

## Architecture

Client → REST API (Gin) → Middleware → Controllers → GORM → PostgreSQL

## Authentication Flow

1. User logs in
2. Password verified using bcrypt
3. JWT token generated
4. Token sent in Authorization header
5. Middleware validates token

## Authorization Flow

- Role extracted from JWT
- Middleware checks role
- Access granted/denied

## Design Decisions

- UUID used for security
- JWT for stateless authentication
- GORM for ORM abstraction
- Swagger for API documentation