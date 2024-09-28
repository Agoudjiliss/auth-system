```markdown
# Go JWT Authentication System

This project implements a simple JWT authentication system in Go, featuring user registration, login, and token management. The application utilizes PostgreSQL for data storage and includes middleware for protecting routes.

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Testing with Postman](#testing-with-postman)
- [License](#license)

## Features

- User registration with password hashing.
- User login with JWT token generation.
- Middleware to protect routes.
- Token expiration handling.
- Database integration with PostgreSQL.
- Refresh token management (optional).

## Technologies Used

- **Go** - Programming Language
- **PostgreSQL** - Database
- **Chi** - HTTP Router
- **Viper** - Configuration Management
- **JWT** - JSON Web Token for authentication
- **bcrypt** - Password hashing

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/auth-system.git
   cd auth-system
   ```

2. Set up your PostgreSQL database:
   - Create a new database for the application.
   - Update the database configuration in the `config.yaml` file.

3. Install Go dependencies:
   ```bash
   go mod tidy
   ```

4. Run the application:
   ```bash
   make run
   ```

5. The server will start on `http://localhost:8080`.

## Usage

### Configuration

Update the `config.yaml` file with your PostgreSQL credentials and JWT secret key:

```yaml
server:
  Host: "localhost"
  Port: "8080"
database:
  User: "your_db_user"
  Password: "your_db_password"
  Dbname: "your_db_name"
  Sslmode: "disable"

jwt:
  Jwtkey: "your_jwt_secret_key"
```

### API Endpoints

- **POST /register**
  - Request body:
    ```json
    {
      "username": "your_username",
      "password": "your_password"
    }
    ```
  - Description: Registers a new user.

- **POST /login**
  - Request body:
    ```json
    {
      "username": "your_username",
      "password": "your_password"
    }
    ```
  - Description: Logs in a user and returns a JWT token as a cookie.

- **GET /protected**
  - Description: A protected route that requires a valid JWT token to access.

## Testing with Postman

1. **Login Request**:
   - **Method**: `POST`
   - **URL**: `http://localhost:8080/connect`
   - **Body**:
     ```json
     {
       "username": "your_username",
       "password": "your_password"
     }
     ```

2. **Protected Endpoint**:
   - **Method**: `GET`
   - **URL**: `http://localhost:8080/hello`
   - Postman will automatically include the cookie set during login.


```

### Notes
- Replace placeholders like `yourusername`, `your_db_user`, and `your_jwt_secret_key` with appropriate values related to your project.
- You can add sections for troubleshooting, contributing, or more detailed examples of each endpoint if needed.
- Consider including a **Contributing** section if you want others to contribute to your project. 

Feel free to ask if you need any changes or additional sections!
