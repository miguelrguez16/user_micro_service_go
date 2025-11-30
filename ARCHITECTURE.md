# User Microservice - Hexagonal Architecture

## Project Structure

The project has been refactored following the **Hexagonal Architecture (Ports & Adapters)** pattern.

```
userMicroservice/
├── cmd/
│   └── main.go                                    # Application entry point
│
├── internal/
│   ├── domain/                                    # Domain Layer (business core)
│   │   ├── models/
│   │   │   └── user.go                           # User Entity
│   │   └── ports/
│   │       ├── repository.go                     # UserRepository Interface
│   │       └── service.go                        # UserService Interface
│   │
│   ├── application/                               # Application Layer (use cases)
│   │   └── usecases/
│   │       └── user_service.go                   # Services / Use Cases
│   │
│   ├── adapters/                                  # Adapters Layer
│   │   ├── http/                                 # HTTP Adapter (inbound)
│   │   │   ├── user_controller.go                # User Controllers
│   │   │   ├── health_controller.go              # Health & Ping Controller
│   │   │   └── server/
│   │   │       └── route.go                      # Route Configuration
│   │   │
│   │   └── persistence/                          # Persistence Adapter (outbound)
│   │       └── mongo/
│   │           └── user_repository.go            # MongoDB Repository Implementation
│   │
│   └── infra/                                     # Infrastructure Layer
│       └── config/
│           └── config.go                         # Configuration & DB Connection
│
├── src/                                           # ⚠️ LEGACY (keep for transition)
│   ├── config/
│   ├── controllers/
│   ├── models/
│   ├── ports/
│   ├── repository/
│   ├── route/
│   └── services/
│
├── go.mod
├── go.sum
└── README.md
```

## Hexagonal Architecture Explained

### 1. **Domain Layer**

- Contains **pure business logic** independent of any framework.
- **models/**: Domain entities (User).
- **ports/**: Interfaces defining contracts that adapters must implement.
  - **UserRepository interface**: Defines operations any repository must implement.

### 2. **Application Layer**

- Contains **use cases** (business services).
- **UserService**: Orchestrates business logic using domain interfaces.
- Depends on **domain** but NOT on adapters.

### 3. **Adapters Layer**

- Connects the domain to the external world (HTTP, Database, APIs, etc.).
- **http/**: HTTP Adapter
  - **UserController**: Handles HTTP requests.
  - **server/route.go**: Configures routes.
- **persistence/mongo/**: Persistence Adapter
  - **UserRepository**: Implements domain interface for MongoDB.

### 4. **Infrastructure Layer**

- Technical configuration: database connection, environment variables, etc.
- **config.go**: MongoDB connection setup.

### 5. **cmd/main.go**

- Application entry point.
- **Orchestrates dependencies** (Dependency Injection):
  1. Connects to database.
  2. Creates MongoDB repository instance.
  3. Creates service instance.
  4. Configures HTTP routes.
  5. Starts the HTTP server.

## Dependency Flow

```
HTTP Request
    ↓
[Adapters HTTP] UserController
    ↓
[Application] UserService ← depends on
    ↓
[Domain Ports] UserRepository (interface)
    ↓
[Adapters Persistence] UserRepository (MongoDB implementation)
    ↓
MongoDB
```

## Hexagonal Architecture Benefits

✅ **Framework Independence**: Domain doesn't depend on Gin, MongoDB, etc.
✅ **Testability**: Easy to mock domain interfaces.
✅ **Scalability**: Easy to add new adapters (e.g., PostgreSQL, external REST API).
✅ **Dependency Inversion**: Dependencies point toward the domain.
✅ **Maintainability**: Well-organized code with clear responsibilities.

## Example: Adding a New Adapter

To switch from MongoDB to PostgreSQL:

1. Create: `internal/adapters/persistence/postgres/user_repository.go`
2. Implement the `ports.UserRepository` interface
3. Update `cmd/main.go` to use the new implementation
4. **Domain and application layers need NO changes**

## Next Steps (Optional)

- [ ] Create unit tests for `UserService`
- [ ] Create integration tests for `UserRepository`
- [ ] Add more use cases (CreateUser, UpdateUser, DeleteUser)
- [ ] Create DTOs (Data Transfer Objects) to separate domain models from API responses
- [ ] Add authentication/authorization middleware
- [ ] Implement custom error handling

## Available Endpoints

### Health & Monitoring

#### Microservice Ping
```
GET /ping
```

Response:
```json
{
  "service": "user-microservice",
  "status": "pong"
}
```

#### Health Check
```
GET /health
```

Response (Healthy - Status 200):
```json
{
  "status": "ok",
  "service": "user-microservice",
  "database": "connected"
}
```

Response (Degraded - Status 503):
```json
{
  "status": "degraded",
  "service": "user-microservice",
  "database": "disconnected"
}
```

Status Code: `503 Service Unavailable` if the database is disconnected.

### User Management

#### Get All Users
```
GET /users
```

#### Check Database Connection
```
GET /users/ping
```

#### Get Total User Count
```
GET /users/total
```
