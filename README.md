## Student CRUD REST API in Golang with Gin
This repository implements a simple REST API for managing student data using Golang and the Gin framework.

## Learning Objectives
* **REST API Best Practices: Follow industry standards for designing and implementing RESTful APIs.
* **Twelve-Factor App Methodology: Adhere to the principles of the Twelve-Factor App methodology for creating portable and maintainable applications.

## Functionality
This API allows you to perform CRUD operations (Create, Read, Update, and Delete) on student data:

* **Create a new student: POST /api/v1/students
* **Get all students: GET /api/v1/students
* **Get a student by ID: GET /api/v1/students/:id
* **Update a student: PUT /api/v1/students/:id
* **Delete a student: DELETE /api/v1/students/:id

PS: the same function and end points are made for courses and enrollment

## Project Setup
1. Clone the repository:
git clone https://github.com/<your_username>/student-api.git
2. Install dependencies:
go mod download
3. Configure environment variables:
Create a .env file in the project root directory and define any necessary environment variables (e.g., database connection details).

Example:

DB_URL=postgres://user:password@localhost:5432/student_db

4. Build and run the API:
make run

## Additional Features
* **API Versioning: Utilizes versioning prefix (/api/v1) for clarity and potential future changes.
* **HTTP Verbs: Employs appropriate HTTP verbs (POST, GET, PUT, DELETE) for specific operations.
* **Logging: Implements informative logs with different log levels for debugging and monitoring.
* **Health Check: Provides a /healthcheck endpoint to verify API health.
* **Unit Tests: Includes unit tests for various API endpoints using the gomock framework.

PS: I'm yet to complete the unit test for the api endpoint. I think i need collaboration on that. I am using gomock to mock the Db and make the test independent of the main DB.

## Resources
Database schema designed using https://app.diagrams.net/
Configuration management with https://github.com/spf13/viper
This project serves as a foundation for building and testing basic CRUD APIs using Golang and Gin. Feel free to explore and extend it further based on your specific requirements.