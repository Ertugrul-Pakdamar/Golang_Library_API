# Library Management API

A modern, RESTful API for library management built with Go and MongoDB. This API provides comprehensive functionality for managing users, books, and borrowing operations in a library system.

## 🚀 Features

- **User Management**: Registration, authentication, and account management
- **Book Management**: Add, list, and manage library catalog
- **Borrowing System**: Borrow and return books with automatic tracking
- **Role-Based Access**: Administrator and member roles with different permissions
- **JWT Authentication**: Secure token-based authentication
- **Clean Architecture**: Well-structured codebase with services, handlers, and middleware
- **RESTful API**: Standard HTTP methods and status codes
- **MongoDB Integration**: NoSQL database for flexible data storage

## 🏗️ Architecture

```
├── handlers/          # HTTP request handlers
├── services/          # Business logic layer
├── middleware/        # Authentication and authorization
├── models/           # Data models
├── database/         # Database connection and configuration
├── utils/            # Utility functions
└── tester/           # Python API tester
```

## 📋 Prerequisites

- Go 1.19 or higher
- MongoDB 4.4 or higher
- Python 3.7+ (for testing)

## 🛠️ Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/Ertugrul-Pakdamar/Golang_Library_API.git
   cd golang_api
   ```

2. **Install dependencies**

   ```bash
   go mod tidy
   ```

3. **Start MongoDB**

   ```bash
   # Using Docker
   docker run -d -p 27017:27017 --name mongodb mongo:latest

   # Or start your local MongoDB instance
   mongod
   ```

4. **Run the application**
   ```bash
   go run main.go
   ```

The API will be available at `http://localhost:3000`

## 📚 API Endpoints

### Authentication

- `POST /api/user/register` - Register a new user
- `POST /api/user/login` - Login and get JWT token
- `DELETE /api/user/delete` - Delete user account (requires authentication)

### User Management

- `GET /api/user/info` - Get user information (requires authentication)

### Book Management

- `POST /api/book/add` - Add book to library (admin only)
- `GET /api/books` - Get all books (requires authentication)

### Borrowing Operations

- `POST /api/book/borrow` - Borrow a book (requires authentication)
- `POST /api/book/return` - Return a book (requires authentication)

## 🔐 Authentication

The API uses JWT (JSON Web Tokens) for authentication. Include the token in the Authorization header:

```
Authorization: Bearer <your-jwt-token>
```

### User Roles

- **Administrator (Role: 0)**: Can add books to the library
- **Member (Role: 1)**: Can borrow and return books

**Note**: The first registered user automatically becomes an administrator.

## 📖 Usage Examples

### Register a User

```bash
curl -X POST http://localhost:3000/api/user/register \
  -H "Content-Type: application/json" \
  -d '{"username": "john_doe", "password": "SecurePass123"}'
```

### Login

```bash
curl -X POST http://localhost:3000/api/user/login \
  -H "Content-Type: application/json" \
  -d '{"username": "john_doe", "password": "SecurePass123"}'
```

### Add a Book (Admin only)

```bash
curl -X POST http://localhost:3000/api/book/add \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-token>" \
  -d '{"title": "The Great Gatsby", "author": "F. Scott Fitzgerald"}'
```

### Borrow a Book

```bash
curl -X POST http://localhost:3000/api/book/borrow \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-token>" \
  -d '{"title": "The Great Gatsby"}'
```

## 🧪 Testing

A Python-based terminal tester is included for easy API testing:

```bash
cd tester
python api_tester.py
```

The tester provides an interactive menu for all API operations.

## 📊 Data Models

### User

```json
{
	"id": "ObjectId",
	"username": "string",
	"password": "string (hashed)",
	"role": "number (0=admin, 1=member)",
	"books_taken": ["ObjectId"]
}
```

### Book

```json
{
	"id": "ObjectId",
	"title": "string",
	"author": "string",
	"count": "number",
	"borrowed": "number"
}
```

## 🔒 Security Features

- **Password Hashing**: Bcrypt encryption for secure password storage
- **JWT Tokens**: Secure authentication with configurable expiration
- **Input Validation**: Comprehensive validation for all inputs
- **Role-Based Access**: Different permissions for different user types
- **Borrowing Limits**: Maximum 2 books per user

## 📝 Business Rules

- Users can borrow a maximum of 2 books at a time
- Only administrators can add books to the library
- Books cannot be borrowed if all copies are already borrowed
- Users cannot borrow the same book twice
- When a user is deleted, their borrowed books are automatically returned

## 🚨 Error Handling

The API returns consistent error responses with appropriate HTTP status codes:

```json
{
	"success": false,
	"message": "Error description",
	"error": {
		"code": 400,
		"details": "Detailed error information"
	}
}
```

## 🛡️ HTTP Status Codes

- `200` - Success
- `400` - Bad Request
- `401` - Unauthorized
- `403` - Forbidden
- `404` - Not Found
- `409` - Conflict
- `422` - Unprocessable Entity
- `500` - Internal Server Error

## 🔧 Configuration

The API uses the following default configuration:

- **Port**: 3000
- **MongoDB**: localhost:27017
- **Database**: golang_api_db
- **Collections**: users, books

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

**Using Go, Fiber, and MongoDB**
