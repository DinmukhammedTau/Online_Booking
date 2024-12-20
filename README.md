# Golang Web Server with JSON Handling and Database Integration

## Project Overview
This project implements a web server in Golang capable of handling HTTP POST and GET requests with JSON data. The server validates JSON fields and responds with appropriate messages. Additionally, the project includes CRUD operations connected to a PostgreSQL database. 

The primary goals of the project are:
- Efficient handling of JSON requests and responses.
- Database integration with migrations and advanced CRUD operations.
- Simple HTML interface for interaction with the backend.

## Features
1. **JSON Request/Response Handling:**
   - Handles POST and GET requests.
   - Validates JSON fields and data types.
   - Returns appropriate success or failure messages in JSON format.

2. **Database Integration:**
   - PostgreSQL setup with migrations.
   - CRUD operations on a `users` table.
   - ORM integration using GORM.

3. **HTML Frontend:**
   - Simple interface for testing server endpoints.
   - Displays server responses in a table.
   - Error handling with pop-up messages for request errors.

## Team Members
- [Name 1](mailto:email@example.com)
- [Name 2](mailto:email@example.com)
- [Name 3](mailto:email@example.com)

## Getting Started
### Prerequisites
- [Go Programming Language](https://golang.org/dl/) (version 1.18 or higher recommended)
- [PostgreSQL](https://www.postgresql.org/)
- [Postman](https://www.postman.com/) (for testing API endpoints)

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/your-repository.git
   cd your-repository
