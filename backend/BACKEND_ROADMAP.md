# 🚀 Lyked Backend Roadmap

## Current State Analysis

### ✅ What's Working

- Go + Gin web framework setup with proper routing
- Dual database architecture: MongoDB for media/folders, PostgreSQL for users
- MongoDB v2 driver properly configured
- Basic upload handlers implemented (Create, Read, Delete)
- CORS configuration for Expo development
- Environment variable loading with utils.GetEnv()
- Modular structure: handlers, routes, models, utilities separated

### 🚧 Current Issues & Limitations

- No authentication middleware - all endpoints are public
- Missing JWT token generation and validation
- Delete endpoint has bug: using string comparison instead of ObjectID
- No folder management endpoints implemented
- Missing input validation and sanitization
- No rate limiting or security headers
- Error responses lack proper structure and logging
- No database connection pooling or optimization
- Missing health checks and monitoring
- No API documentation (OpenAPI/Swagger)

---

## 🎯 Development Roadmap

### Phase 1: Authentication & Security

**Priority: Critical** | **Estimated Time: 1-2 weeks**

#### 1.1 JWT Authentication System

- [ ] **JWT Token Management**

  - [ ] Install JWT library: `go get github.com/golang-jwt/jwt/v5`
  - [ ] Create JWT utility functions:
    - [ ] `GenerateToken(userID string) (string, error)`
    - [ ] `ValidateToken(tokenString string) (*Claims, error)`
    - [ ] `RefreshToken(tokenString string) (string, error)`

- [ ] **Authentication Middleware**
  - [ ] Create `AuthMiddleware()` function
  - [ ] Extract and validate JWT from Authorization header
  - [ ] Add user context to Gin context for downstream handlers
  - [ ] Handle expired tokens with proper error responses

**Implementation Details:**

```go
// middleware/jwt.go
type Claims struct {
    UserID string `json:"user_id"`
    jwt.RegisteredClaims
}

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Extract token from Authorization header
        // Validate token and extract claims
        // Set user context for handlers
        // Handle errors with proper HTTP responses
    }
}
```

#### 1.2 User Authentication Endpoints

- [ ] **User Registration & Login**

  - [ ] `POST /auth/register` - Create new user account
  - [ ] `POST /auth/login` - Authenticate user credentials
  - [ ] `POST /auth/refresh` - Refresh expired JWT token
  - [ ] `POST /auth/logout` - Invalidate token (blacklist)

- [ ] **Password Management**

  - [ ] Hash passwords using bcrypt: `go get golang.org/x/crypto/bcrypt`
  - [ ] `POST /auth/forgot-password` - Send password reset email
  - [ ] `POST /auth/reset-password` - Reset password with token
  - [ ] `PUT /auth/change-password` - Change password (authenticated)

- [ ] **User Profile Management**
  - [ ] `GET /users/profile` - Get current user profile
  - [ ] `PUT /users/profile` - Update user profile
  - [ ] `DELETE /users/account` - Delete user account and data

#### 1.3 Security Hardening

- [ ] **Input Validation & Sanitization**

  - [ ] Install validation library: `go get github.com/go-playground/validator/v10`
  - [ ] Create validation structs for all request bodies
  - [ ] Implement custom validation rules for URLs, tags, etc.
  - [ ] Sanitize user input to prevent XSS/injection attacks

- [ ] **Security Headers & Middleware**

  - [ ] Add security headers (HSTS, CSP, X-Frame-Options)
  - [ ] Implement rate limiting: `go get github.com/gin-contrib/timeout`
  - [ ] Add request size limits
  - [ ] Implement CSRF protection for web clients

- [ ] **Database Security**
  - [ ] Use parameterized queries to prevent injection
  - [ ] Implement proper error handling without data leakage
  - [ ] Add database connection encryption (SSL/TLS)
  - [ ] Implement query timeout and connection limits

---

### Phase 2: Core API Development

**Priority: Critical** | **Estimated Time: 2-3 weeks**

#### 2.1 Complete Upload Management

- [ ] **Fix Existing Upload Handlers**

  - [ ] Fix DeleteUploadHandler ObjectID conversion bug:
    ```go
    // Current (broken):
    collection.DeleteOne(ctx, primitive.M{"_id": id})
    // Fixed:
    objectID, _ := primitive.ObjectIDFromHex(id)
    collection.DeleteOne(ctx, primitive.M{"_id": objectID})
    ```

- [ ] **Enhanced Upload Validation**

  - [ ] Validate URL format and accessibility
  - [ ] Check for duplicate URLs per user
  - [ ] Validate title length and content
  - [ ] Sanitize description and tag inputs
  - [ ] Implement file size limits for media URLs

- [ ] **Upload Metadata Enhancement**
  - [ ] Add creation/modification timestamps
  - [ ] Implement user ownership validation
  - [ ] Add view count tracking
  - [ ] Support for custom thumbnail URLs
  - [ ] Add media type detection (image/video/link)

#### 2.2 Folder Management System

- [ ] **Folder CRUD Operations**

  - [ ] `POST /folders` - Create new folder
  - [ ] `GET /folders` - List user's folders
  - [ ] `GET /folders/:id` - Get folder details with contents
  - [ ] `PUT /folders/:id` - Update folder details
  - [ ] `DELETE /folders/:id` - Delete folder (move contents to root)

- [ ] **Folder Organization Features**
  - [ ] `POST /folders/:id/uploads` - Add uploads to folder
  - [ ] `DELETE /folders/:id/uploads/:uploadId` - Remove upload from folder
  - [ ] `PUT /folders/reorder` - Reorder folders
  - [ ] Support for nested folders (max 2 levels)

**Implementation Details:**

```go
// model/mongoModels/folders.go
type Folder struct {
    ID          primitive.ObjectID `bson:"_id" json:"id"`
    UserID      string            `bson:"user_id" json:"user_id"`
    Name        string            `bson:"name" json:"name"`
    Description string            `bson:"description" json:"description"`
    ParentID    *primitive.ObjectID `bson:"parent_id,omitempty" json:"parent_id,omitempty"`
    Order       int               `bson:"order" json:"order"`
    CreatedAt   time.Time         `bson:"created_at" json:"created_at"`
    UpdatedAt   time.Time         `bson:"updated_at" json:"updated_at"`
}
```

#### 2.3 Search & Filtering APIs

- [ ] **Advanced Search Endpoint**

  - [ ] `GET /search` - Full-text search across uploads
  - [ ] Support search by: title, description, tags, URL domain
  - [ ] Implement search filters: date range, folder, media type
  - [ ] Add search result ranking and relevance scoring
  - [ ] Return search suggestions and autocomplete

- [ ] **Filtering & Sorting**
  - [ ] `GET /uploads/filter` - Advanced filtering options
  - [ ] Filter by tags (AND/OR logic)
  - [ ] Filter by folder membership
  - [ ] Filter by date created/modified
  - [ ] Sort by: date, title, view count, relevance

#### 2.4 Tags & Categories System

- [ ] **Tag Management**

  - [ ] `GET /tags` - Get user's most used tags
  - [ ] `POST /tags/suggestions` - Get tag suggestions based on URL content
  - [ ] `PUT /uploads/:id/tags` - Update upload tags
  - [ ] Implement tag autocomplete and validation

- [ ] **Category Analytics**
  - [ ] `GET /analytics/tags` - Tag usage statistics
  - [ ] `GET /analytics/folders` - Folder statistics
  - [ ] `GET /analytics/activity` - User activity over time

---

### Phase 3: Database & Performance Optimization

**Priority: Medium** | **Estimated Time: 1-2 weeks**

#### 3.1 Database Indexing & Query Optimization

- [ ] **MongoDB Indexes**

  - [ ] Create indexes on frequently queried fields:
    ```go
    // user_id index for fast user queries
    // text index for full-text search
    // compound indexes for common filter combinations
    ```

- [ ] **Query Optimization**
  - [ ] Implement pagination with proper cursor-based navigation
  - [ ] Use MongoDB aggregation pipeline for complex queries
  - [ ] Implement query result caching with Redis
  - [ ] Add query execution time monitoring

#### 3.2 Connection Pooling & Database Management

- [ ] **Connection Optimization**

  - [ ] Configure MongoDB connection pool settings
  - [ ] Implement PostgreSQL connection pooling
  - [ ] Add connection health checks and auto-retry logic
  - [ ] Monitor and log database connection metrics

- [ ] **Data Migration & Backup**
  - [ ] Create database migration system
  - [ ] Implement automated backup procedures
  - [ ] Add data cleanup jobs for deleted users
  - [ ] Create data export functionality for users

#### 3.3 Caching Strategy

- [ ] **Redis Integration**

  - [ ] Install Redis client: `go get github.com/redis/go-redis/v9`
  - [ ] Cache frequently accessed user data
  - [ ] Cache search results and tag suggestions
  - [ ] Implement cache invalidation strategies
  - [ ] Add cache warming for popular content

- [ ] **Application-Level Caching**
  - [ ] Cache LinkPreview API responses
  - [ ] Implement in-memory caching for configuration
  - [ ] Add ETag support for conditional requests

---

### Phase 4: Production Readiness

**Priority: Medium** | **Estimated Time: 1-2 weeks**

#### 4.1 Comprehensive Logging & Monitoring

- [ ] **Structured Logging**

  - [ ] Replace fmt.Printf with proper logging: `go get github.com/sirupsen/logrus`
  - [ ] Implement log levels: DEBUG, INFO, WARN, ERROR
  - [ ] Add request ID tracking for correlation
  - [ ] Log API response times and error rates

- [ ] **Metrics & Monitoring**
  - [ ] Add Prometheus metrics: `go get github.com/prometheus/client_golang`
  - [ ] Monitor API endpoint performance
  - [ ] Track database query performance
  - [ ] Monitor memory and CPU usage

#### 4.2 Error Handling & Recovery

- [ ] **Centralized Error Handling**

  - [ ] Create custom error types with error codes
  - [ ] Implement error wrapping with context
  - [ ] Add error recovery middleware
  - [ ] Create user-friendly error messages

- [ ] **Graceful Shutdown**
  - [ ] Implement graceful server shutdown
  - [ ] Close database connections properly
  - [ ] Complete in-flight requests before shutdown
  - [ ] Add health check endpoint for load balancers

#### 4.3 API Documentation & Testing

- [ ] **OpenAPI/Swagger Documentation**

  - [ ] Install Swagger: `go get github.com/swaggo/gin-swagger`
  - [ ] Document all API endpoints with examples
  - [ ] Include authentication requirements
  - [ ] Add response schema definitions

- [ ] **Comprehensive Testing**
  - [ ] Unit tests for all handlers and utilities
  - [ ] Integration tests for database operations
  - [ ] API endpoint tests with test database
  - [ ] Load testing with realistic traffic patterns

---

## 🔧 Environment & Configuration

### Required Environment Variables

Create a `.env` file in the backend/ directory:

```bash
# Server Configuration
PORT=3000
GIN_MODE=release
CORS_ORIGINS=http://localhost:8081,exp://192.168.*:8081

# Database Configuration
MONGODB_URI=mongodb://localhost:27017/lyked
POSTGRES_URI=postgres://user:password@localhost:5432/lyked

# Authentication
JWT_SECRET_KEY=your-super-secret-jwt-key-here
JWT_EXPIRY_HOURS=24
REFRESH_TOKEN_EXPIRY_DAYS=7

# External APIs
LINKPREVIEW_API_KEY=your-linkpreview-api-key
LINKPREVIEW_RATE_LIMIT=100

# Redis Configuration (optional)
REDIS_URI=redis://localhost:6379
REDIS_DB=0

# Security
BCRYPT_COST=12
RATE_LIMIT_REQUESTS_PER_MINUTE=60

# Monitoring
ENABLE_METRICS=true
LOG_LEVEL=INFO
```

### Required Dependencies to Add

```bash
# Authentication & Security
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
go get github.com/gin-contrib/cors

# Database & Caching
go get github.com/redis/go-redis/v9
go get gorm.io/driver/postgres

# Validation & Utilities
go get github.com/go-playground/validator/v10
go get github.com/google/uuid

# Logging & Monitoring
go get github.com/sirupsen/logrus
go get github.com/prometheus/client_golang

# Documentation
go get github.com/swaggo/gin-swagger
go get github.com/swaggo/files

# Development & Testing
go get github.com/stretchr/testify
go get github.com/DATA-DOG/go-sqlmock
```

---

## 🏗️ Architecture Improvements

### Current Structure Enhancement

```
backend/
├── cmd/
│   └── main.go
├── internal/                    # New: Internal packages
│   ├── auth/                   # Authentication logic
│   ├── config/                 # Configuration management
│   ├── middleware/             # All middleware
│   └── services/               # Business logic layer
├── api/                        # New: API layer
│   ├── handlers/              # HTTP handlers (moved)
│   ├── routes/                # Route definitions
│   └── validators/            # Request validation
├── pkg/                       # New: Reusable packages
│   ├── database/              # Database utilities
│   ├── logger/                # Logging utilities
│   └── errors/                # Custom error types
└── tests/                     # New: Test files
    ├── integration/
    └── unit/
```

### Recommended Code Patterns

#### 1. Service Layer Pattern

```go
// internal/services/upload_service.go
type UploadService struct {
    mongoRepo *repositories.MongoUploadRepository
    validator *validator.Validate
}

func (s *UploadService) CreateUpload(ctx context.Context, userID string, req CreateUploadRequest) (*model.LykedUploads, error) {
    // Business logic here
    // Validation, data processing, etc.
}
```

#### 2. Repository Pattern

```go
// internal/repositories/upload_repository.go
type UploadRepository interface {
    Create(ctx context.Context, upload *model.LykedUploads) error
    GetByUserID(ctx context.Context, userID string, opts GetOptions) ([]*model.LykedUploads, error)
    Delete(ctx context.Context, id primitive.ObjectID) error
}
```

---

## 📋 API Endpoint Specifications

### Authentication Endpoints

```
POST   /api/v1/auth/register       # User registration
POST   /api/v1/auth/login          # User login
POST   /api/v1/auth/refresh        # Refresh token
POST   /api/v1/auth/logout         # Logout user
POST   /api/v1/auth/forgot-password # Password reset request
POST   /api/v1/auth/reset-password  # Password reset confirmation
```

### Upload Management

```
POST   /api/v1/uploads             # Create upload
GET    /api/v1/uploads             # Get user uploads (paginated)
GET    /api/v1/uploads/:id         # Get single upload
PUT    /api/v1/uploads/:id         # Update upload
DELETE /api/v1/uploads/:id         # Delete upload
POST   /api/v1/uploads/bulk        # Bulk operations
```

### Folder Management

```
POST   /api/v1/folders             # Create folder
GET    /api/v1/folders             # Get user folders
GET    /api/v1/folders/:id         # Get folder with contents
PUT    /api/v1/folders/:id         # Update folder
DELETE /api/v1/folders/:id         # Delete folder
POST   /api/v1/folders/:id/uploads # Add uploads to folder
```

### Search & Analytics

```
GET    /api/v1/search              # Search uploads
GET    /api/v1/tags                # Get user tags
GET    /api/v1/analytics/summary   # User analytics
GET    /api/v1/health              # Health check
```

---

## 🚨 Security Checklist

- [ ] **Authentication**: JWT tokens with proper expiration
- [ ] **Authorization**: User can only access their own data
- [ ] **Input Validation**: All inputs validated and sanitized
- [ ] **SQL/NoSQL Injection**: Parameterized queries only
- [ ] **Rate Limiting**: Prevent API abuse
- [ ] **HTTPS**: All traffic encrypted in production
- [ ] **CORS**: Properly configured origins
- [ ] **Headers**: Security headers implemented
- [ ] **Secrets**: No hardcoded secrets or API keys
- [ ] **Error Handling**: No sensitive data in error messages

---

## 📊 Performance Targets

- [ ] **Response Time**: < 200ms for 95% of requests
- [ ] **Throughput**: Handle 1000+ requests per minute
- [ ] **Database Queries**: < 100ms for 95% of queries
- [ ] **Memory Usage**: < 512MB under normal load
- [ ] **Error Rate**: < 1% of requests fail
- [ ] **Uptime**: 99.9% availability target

---

## 🧪 Testing Strategy

### Unit Testing (70% coverage minimum)

- [ ] Test all handler functions with mock dependencies
- [ ] Test service layer business logic
- [ ] Test utility functions and validators
- [ ] Test database repositories with test database

### Integration Testing

- [ ] Test complete API workflows
- [ ] Test database migrations and operations
- [ ] Test authentication flows
- [ ] Test error scenarios and edge cases

### Load Testing

- [ ] Test with realistic user traffic patterns
- [ ] Test database performance under load
- [ ] Test memory and CPU usage under stress
- [ ] Test graceful degradation scenarios

---

## 📋 Definition of Done

For each backend feature to be considered complete, it must:

1. **✅ Functionality**: Feature works as designed with proper error handling
2. **✅ Security**: Authentication/authorization implemented correctly
3. **✅ Performance**: Meets response time and throughput targets
4. **✅ Testing**: Unit and integration tests written and passing
5. **✅ Documentation**: API endpoints documented with examples
6. **✅ Logging**: Proper logging and monitoring implemented
7. **✅ Code Review**: Code follows Go conventions and passes linting

---

## 🚀 Deployment Checklist

Before production deployment:

- [ ] **Environment**: All environment variables configured
- [ ] **Database**: Migrations applied and indexes created
- [ ] **Monitoring**: Logging and metrics collection active
- [ ] **Security**: Security audit completed
- [ ] **Performance**: Load testing passed
- [ ] **Documentation**: API documentation published
- [ ] **Backup**: Database backup procedures tested
- [ ] **Recovery**: Disaster recovery plan documented

---

## 📊 Progress Tracking

Track your progress by checking off completed items. Focus on completing Phase 1 (Authentication & Security) before moving to other phases.

**Current Phase**: Phase 1 - Authentication & Security  
**Next Milestone**: JWT authentication middleware implemented

---

_This roadmap represents a production-ready backend architecture. Adjust timelines and priorities based on your specific requirements and constraints._
