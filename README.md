## Student CRUD REST API in Golang with Gin (following One2n Playbook SRE Challenge)
This repository implements a simple REST API for managing student data using Golang and the Gin framework.

### Learning Objectives
* REST API Best Practices: Follow industry standards for designing and implementing RESTful APIs.
* Twelve-Factor App Methodology: Adhere to the principles of the Twelve-Factor App methodology for creating portable and maintainable applications.

### Functionality
This API allows you to perform CRUD operations (Create, Read, Update, and Delete) on student data:

* Create a new student: POST /api/v1/students
* Get all students: GET /api/v1/students
* Get a student by ID: GET /api/v1/students/:id
* Update a student: PUT /api/v1/students/:id
* Delete a student: DELETE /api/v1/students/:id

PS: the same function and end points are made for courses and enrollment

### Project Setup
1. Clone the repository:
git clone https://github.com/<your_username>/student-api.git
2. Install dependencies:
go mod download
3. Configure environment variables:
Create app.env file in the project root directory and define any necessary environment variables (e.g., database connection details).

Example:

DB_URL=postgres://user:password@localhost:5432/student_db

4. Build and run the API:
make run

### Additional Features
* API Versioning: Utilizes versioning prefix (/api/v1) for clarity and potential future changes.
* HTTP Verbs: Employs appropriate HTTP verbs (POST, GET, PUT, DELETE) for specific operations.
* Logging: Implements informative logs with different log levels for debugging and monitoring.
* Health Check: Provides a /healthcheck endpoint to verify API health.
* Unit Tests: Includes unit tests for various API endpoints using the gomock framework.

PS: I'm yet to complete the unit test for the api endpoint. I think i need collaboration on that. I am using gomock to mock the Db and make the test independent of the main DB.

### Resources
Database schema designed using https://app.diagrams.net/
Configuration management with https://github.com/spf13/viper
This project serves as a foundation for building and testing basic CRUD APIs using Golang and Gin. Feel free to explore and extend it further based on your specific requirements.

### 2 - Containerise REST API (Building and Running a Docker Container)
Docker is used to containerized the API application. To build the image using docker:
* Building the Docker Image
1. Clone this repository to your local machine:
git clone https://github.com/your_username/project_name.git
2. Navigate to the project directory:
cd project_name
3. Build the Docker image using the provided Dockerfile:
docker build -t project_name .

##### PS: Replace project_name with your desired image name.

* Running the Docker Container
1. Once the Docker image is built, you can run a container using the following command:
docker run --name container_name -p host_port:container_port project_name

Replace:

* container_name with your desired container name.
* host_port with the port on your host machine that you want to map to the container.
* container_port with the port inside the container where your application is listening.
* project_name with the name of the Docker image you built earlier.

2. Access your application in a web browser or through API calls using:
* localhost:host_port.

### To stop and remove a running container, use:
* docker rm -f container_name
To remove the Docker image, use:
* docker rmi project_name

### Milestone 3 --> Setup one-click local development setup

On the basis that the team is familiar with the development environment and have all the required tools (Docker, Docker Compose, Make) on their locall machine,depending on the OS each team member is using. This milestone simplify the process of setting up the API on a local machine.
The process of developing and deploying the project have been streamlined using docker compose for defining and running the API and database container. the setup is done the right order where the database is start and ready before the API container is spinned up.

Upon cloning the project repo to your local make, run the following code/script to start the application:

--> make start_db
--> make run_api

To top the containers:
--> make stop_api

### Milestone 4 --> Setup a CI pipeline

Resources: 
- https://docs.github.com/en/actions/learn-github-actions/understanding-github-actions#create-an-example-workflow

The main focus of this module is to create a simple CI pipeline that will be used to build and push the project docker image to a central registry (I am using dockerhub).