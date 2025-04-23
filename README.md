# Basic Go Authentication System

This project is a basic authentication system built using the Go programming language. It provides essential features such as user registration, login, role-based access control, and token-based authentication using JSON Web Tokens (JWT). The application is designed to demonstrate how to implement secure authentication and authorization in a Go web application.

## Features

- **User Registration**: Allows users to register with unique email addresses and hashed passwords.
- **User Login**: Authenticates users and generates JWT tokens for session management.
- **Role-Based Access Control**: Restricts access to specific routes based on user roles (e.g., "user" or "admin").
- **Token-Based Authentication**: Uses JWT for secure and stateless authentication.
- **Logout**: Invalidates the user's session by clearing the JWT token.

## Tech Stack

### Backend
- **Go**: The primary programming language used for building the application.
- **Gin**: A lightweight and fast web framework for handling HTTP requests and routing.
- **GORM**: An ORM library for interacting with the PostgreSQL database.
- **PostgreSQL**: The relational database used for storing user data.
- **JWT (github.com/dgrijalva/jwt-go)**: A library for generating and validating JSON Web Tokens.
- **bcrypt (golang.org/x/crypto/bcrypt)**: Used for securely hashing and comparing passwords.
- **godotenv**: A library for loading environment variables from a `.env` file.

### Environment Configuration
The application uses a `.env` file to store sensitive configuration details such as database credentials and the JWT secret key.

### Project Structure
- **`main.go`**: The entry point of the application. Initializes the database, loads environment variables, and sets up routes.
- **`models/`**: Contains database models and configuration for initializing the database.
- **`controllers/`**: Implements the core logic for authentication and authorization.
- **`routes/`**: Defines the API endpoints and maps them to the corresponding controller functions.
- **`utils/`**: Provides utility functions for password hashing, token parsing, and comparison.

### Security
- Passwords are hashed using bcrypt before being stored in the database.
- JWT tokens are signed using a secret key stored in the `.env` file.
- Role-based access control ensures that only authorized users can access specific routes.

## How to Run

1. Clone the repository.
2. Create a `.env` file with the required environment variables (refer to the provided `.env` file).
3. Install dependencies using `go mod tidy`.
4. Run the application using `go run main.go`.
5. Access the application at `http://localhost:8080`.

This project serves as a foundation for building more complex authentication systems and can be extended with additional features such as email verification, password reset, and OAuth integration.