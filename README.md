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
```
git clone https://github.com/<your_username>/student-api.git
```
2. Install dependencies:
```
go mod download
```
3. Configure environment variables:
Create app.env file in the project root directory and define any necessary environment variables (e.g., database connection details).

Example:

*DB_URL=postgres://user:password@localhost:5432/student_db*

4. Build and run the API:
```
make run
```
### Additional Features
* API Versioning: Utilizes versioning prefix (/api/v1) for clarity and potential future changes.
* HTTP Verbs: Employs appropriate HTTP verbs (POST, GET, PUT, DELETE) for specific operations.
* Logging: Implements informative logs with different log levels for debugging and monitoring.
* Health Check: Provides a /healthcheck endpoint to verify API health.
* Unit Tests: Includes unit tests for various API endpoints using the gomock framework.

**PS: I'm yet to complete the unit test for the api endpoint. I think i need collaboration on that. I am using gomock to mock the Db and make the test independent of the main DB.**

### Resources
Database schema designed using [https://app.diagrams.net/]
Configuration management with [https://github.com/spf13/viper]
This project serves as a foundation for building and testing basic CRUD APIs using Golang and Gin. Feel free to explore and extend it further based on your specific requirements.

### 2 - Containerise REST API (Building and Running a Docker Container)
Docker is used to containerized the API application. To build the image using docker:
* Building the Docker Image
1. Clone this repository to your local machine:
```
git clone https://github.com/your_username/project_name.git
```
2. Navigate to the project directory:
```
cd project_name
```
3. Build the Docker image using the provided Dockerfile:
```
docker build -t project_name .
```

##### PS: Replace project_name with your desired image name.

* Running the Docker Container
1. Once the Docker image is built, you can run a container using the following command:
```
docker run --name container_name -p host_port:container_port project_name
```
Replace:

* container_name with your desired container name.
* host_port with the port on your host machine that you want to map to the container.
* container_port with the port inside the container where your application is listening.
* project_name with the name of the Docker image you built earlier.

2. Access your application in a web browser or through API calls using:
```
localhost:host_port.
```

### To stop and remove a running container, use:
```
docker rm -f container_name
```
To remove the Docker image, use:
```
docker rmi project_name
```
## Milestone 3 --> Setup one-click local development setup

On the basis that the team is familiar with the development environment and have all the required tools (Docker, Docker Compose, Make) on their locall machine,depending on the OS each team member is using. This milestone simplify the process of setting up the API on a local machine.
The process of developing and deploying the project have been streamlined using docker compose for defining and running the API and database container. the setup is done the right order where the database is start and ready before the API container is spinned up.

Upon cloning the project repo to your local make, run the following code/script to start the application:
```
make start_db
make run_api
```
To top the containers:
```
make stop_api
```
## Milestone 4 --> Setup a CI pipeline

Resources: 
- https://docs.github.com/en/actions/learn-github-actions/understanding-github-actions#create-an-example-workflow

The main focus of this module is to create a simple CI pipeline that will be used to build and push the project docker image to a central registry (I am using dockerhub registry).


After lauching the ec2 instance, for this project some dependencies needs to install to prepare the environment for runnig the github self-hosted runner. i will be installing docker and git on the instance

PS: Ensure to give the neccessary permission on the ec2 instance to be able to connect and comminicate with github and docker. update the security group setting to http and https traffic.

copy and run the following code line by line to install docker:
```
sudo apt update
curl -fsSL https://get.docker.com -o get-docker.sh
sudo chmod +x get-docker.sh
sudo ./get-docker.sh
sudo docker --version
```
give the user the neccessary permission to e able to run docker command
```
sudo usermod -aG docker $USER
```
```
Logout and log back in again
```

Optional: check docker command path
```
which docker
```
response ==> "/usr/bin/docker"

Install Git using the following command:
```
sudo apt install -y git
git --version
```

After preparing the environment for the self-hosted runner. go to the project repo on github to get and run the runner code on the machine youu want to use for the self-hosted runner. it is found in the setting tab -> Action -> Runner then click on New-self-hosted-runner.

PS: Before running the running the runner ./run.sh command, make sure to logout and login again into the machine to validate the docker permission.

## Milestone 5 --> Deploy REST API & its dependent services on bare metal
![vagrant-deployment](https://github.com/juw0n/SRE-Devop-Bootcamp/assets/45376257/1cdb15aa-1cdc-4485-92d6-37a897d18dbc)

In this milestone, I deployed the REST API & its dependent services on bare metal using vagrant to provisioned a Virtual machine(VM) running Ubuntu on Oracle VM Virtualbox.
I automated the provisioning of the Virtual Machine(VM) using vagrantfile that also install the basic dependencies packages for the project. For setting up the Vagrant box I used a bash script with functions to install the required dependencies. I updated the docker-compose file and Makefile to do the deployment and the final setup should consist of:
* 2 API containers
* 1 DB container
* 1 Nginx container

as shown in the diagram above.
The Nginx was used for load balancing between these two API containers. Internally the Nginx was set to load balance requests between two API containers and is accessible via port 8084. (I had issues accessing the api on my local machine from port 8080, it doesn't matter though).

Resources:
* [Vagrant Documentation](https://developer.hashicorp.com/vagrant/docs)
* [Deploying Nginx using Docker](https://docs.nginx.com/nginx/admin-guide/installing-nginx/installing-nginx-docker/)

## Milestone 6 --> Setup Kubernetes cluster

The expectation of this milestone is to setup a 3 node kubernetes cluster using Minikube. the nodes are then to be appropriately lablled as:
=> Node A: type=application
=> Node B: type=dependent_services
=> Node C: type=observability.

After studing [Minikube Docs](https://minikube.sigs.k8s.io/docs/start/) and following the getting-started-guide, I was able to install minikube and kubectl to manage the cluster. I also learn the basic of kubernete and its' commands also following some parts of it's [Documentation](https://kubernetes.io/docs/home/).

Some highlight of the learnt commands that i used:
1. Starting the Cluster:
==> minikube start --driver=kvm2 --nodes 3 -p sre-project
2. Verify cluster status:
==> minikube status -p sre-project
3. List nodes:
==> kubectl get node -o wide
==> kubectl get nodes --show-labels
4. Label worker nodes:
* Label application node
kubectl label node <node-name> application=application-node
* Label database node
kubectl label node <node-name> database=database-node
* Label dependent-services node
kubectl label node <node-name> service=dependent-services
etc

## Milestone 7 --> Deploy REST API & its dependent services in K8s
