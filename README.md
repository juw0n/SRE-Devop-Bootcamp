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
To stop the containers:
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

## Milestone 6 --> Setup 3 nodes Kubernetes cluster

### Problem Statement
We need to spin up a three-node Kubernetes cluster using Minikube on your local. Going forward we will treat this Kubernetes cluster as our production cluster. 
Out of these three nodes
** Node A will be used by our application.
** Node B will be used for running dependent services.
** Node C will be used for running our observability stack.

1. Starting the Cluster:
==> minikube start --driver=kvm2 --nodes 3 -p sre-project
2. Verify cluster status:
==> minikube status -p sre-project
==> minikube dashboard -p sre-project
==> kubectl get node -o wide
==> kubectl get nodes --show-labels
==> kubectl describe nodes
3. Label the Nodes:
Use a bash or ansible script or run it manually as follows
==> kubectl label node sre-project type=dependent-services-node
==> kubectl label node sre-project-m02 type=database-node
==> kubectl label node sre-project-m03 type=application-node
and
==> kubectl label node sre-project application=application-node
==> kubectl label node sre-project-m02 database=database-node
==> kubectl label node sre-project-m03 application=application-node

Verify the labels have been applied successfully:
==> kubectl label --list nodes node_name
==> kubectl get nodes -l type
==> kubectl get nodes --show-labels

overwrite the node label
==> kubectl label --overwrite nodes node-name key=new-value
To remove the label from a node, provide the key without any value.
==> kubectl label --overwrite nodes node-name key-

OR Taint the nodes:
==> kubectl taint nodes sre-project service=dependent-services-node:NoSchedule
==> kubectl taint nodes sre-project-m02 database=database-node:NoSchedule
==> kubectl taint nodes sre-project-m03 application=application-node:NoSchedule

verify that taint was successfully applied:
==> kubectl describe node <node-name>

Stop running cluster
==> minikube stop -p <profile-name>
==> minikube stop -p sre-project

Delete the cluster
==> minikube delete -p <profile-name>
==> minikube delete -p sre-project

After studing [Minikube Docs](https://minikube.sigs.k8s.io/docs/start/) and following the getting-started-guide, I was able to install minikube and kubectl to manage the cluster. I also learn the basic of kubernete and its' commands also following some parts of it's [Documentation](https://kubernetes.io/docs/home/).

## Milestone 7 --> Deploy REST API & its dependent services in K8s

Step 1 -> Start the cluster, label and taint the nodes appropriately (minikube)
Step 2 -> Create the 3 required namespaces (student-api-ns, vault-ns, external-secrets-ns)
==> kubectl get namespaces --show-labels
Step 3 -> Install Helm chart
Step 4 -> Install the Vault Helm chart
==> helm repo add hashicorp https://helm.releases.hashicorp.com
==> helm repo update
==> helm repo list
==> helm search repo hashicorp/vault
==> helm install vault hashicorp/vault --namespace vault-ns -f vault-values.yaml
Step 5: Configure Vault
-- First, get the Vault Pod name: 
==> kubectl get pods --namespace vault-ns
-- Then, initialize Vault: 
==> kubectl exec -ti <pod-name> -n <namespace> -- vault operator init
==> kubectl exec -ti vault-0 -n vault-ns -- vault operator init
** This command will output several unseal keys and a root token. Save these securely as you will need them to unseal the Vault and for administrative tasks. **
then:
    Unseal Vault with at least three of the five keys:
    => kubectl exec -ti <vault-pod-name> -n <vault-namespace> -- vault operator unseal <unseal-key-1>
==> kubectl exec -ti vault-0 -n vault-ns -- vault operator unseal <Unseal-key>
-> until you get output that says "initialized: true & Sealed: false"

Join the vault-1 and vault-2 pod to the Raft cluster.
==> kubectl exec -ti vault-1 -n vault-ns -- vault operator raft join http://vault-0.vault-internal:8200

==> kubectl exec -ti vault-2 -n vault-ns -- vault operator raft join http://vault-0.vault-internal:8200

Use the unseal key from above to unseal vault-1 and vault-2
==> kubectl exec -ti vault-1 -- vault operator unseal $VAULT_UNSEAL_KEY
==> kubectl exec -ti vault-2 -- vault operator unseal $VAULT_UNSEAL_KEY

-- Set a secret in Vault
1. Access Vault
==> kubectl exec --stdin=true --tty=true vault-0 -n vault-ns -- /bin/sh
2. Login with the root token when prompted.
==> vault login
3. Enable an instance of the kv-v2 secrets engine at the path <specify path> e.g secret
==> vault secrets enable -path=secretdata kv-v2
==> vault secrets enable -path=apidata kv-v2
4. Create a secret at path "secret/" with a username and password.
==> vault kv put secretdata/postgresql POSTGRES_DB="db_name" POSTGRES_USER="db_user" POSTGRES_PASSWORD="db_password"

==> vault kv put apidata/studentapi DB_DRIVER="db_driver_name" DB_SOURCE="db_source_credentials"

5. Verify that the secret is defined at the path "specifiedPath"
==> vault kv get secretdata/postgresql
==> vault kv get apidata/studentapi

-- Configure Kubernetes authentication
==> vault auth enable kubernetes
#### This method did not work for me at this time
a. Configure the Kubernetes authentication method to use the location of the Kubernetes API.
==> vault write auth/kubernetes/config \
  token_reviewer_jwt="$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)" \
  kubernetes_host="https://$(kubectl get svc -n kube-system kubernetes -o jsonpath='{.spec.clusterIP}'):443" \
  kubernetes_ca_cert=@/var/run/secrets/kubernetes.io/serviceaccount/ca.crt

b. Define the policy rules to allow the ESO to read and retrieve secrets stored in Vault at path secretdata/postgresql:

==> vault policy write eso-policy -<<EOF
path "secretdata/postgresql" {
  capabilities = ["read"]
}
EOF
c. Create a Kubernetes authentication role (eso-role) that connects the Kubernetes service account name and eso-policy policy.
==> vault write auth/kubernetes/role/eso-role \
    bound_service_account_names=vault \
    bound_service_account_namespaces=vault-ns \
    policies=eso-policy \
    ttl=24h
d. Check the service name of the role
==> vault read auth/kubernetes/role/eso-role
#### This is the method that work for me at this time
Create Generic token: Create the Vault token secret to be used in the created SecretStore
kubectl create secret generic vault-token --from-literal=token=<YOUR_VAULT_TOKEN> -n <namespace>
==> kubectl create secret generic vault-token -n student-api-ns --from-literal=token=<VAULT_TOKEN>

6. Exit the vault-0 pod.
==> exit

==> kubectl get secrets -n student-api-ns
==> kubectl get secret vault-token -n student-api-ns -o yaml

Step 6: Configure External Secrets Operator (ESO)
## Deploy External Secrets Operator (ESO)
1. add the ESO Helm repository.
==> helm repo add external-secrets https://charts.external-secrets.io
==> helm repo update
==> helm search repo external-secrets/external-secrets
#### Prepare Custom Values File for scheduling the eso on a specific node. eso-node-values.yaml

2. Install ESO Using Helm
==> helm install external-secrets external-secrets/external-secrets \
  --namespace external-secrets-ns \
  -f eso-node-values.yaml

-> Create the ESO custom resource definitions (SecretStore or ClusterSecretStore and ExternalSecreteStore) in the same namespace where the token was created.
***First get and update the vault server address in the secretstote file before applying
==> kubectl get svc -n <vault-namespace>
Then apply the eso CRD's:
==> kubectl apply -f secret-store.yaml
** Check the secretstore resource to confirm it is in ready Statement
==> kubectl get secretstore -n student-api-ns
==> kubectl describe secretstore vault-backend -n student-api-ns
-> kubectl delete secretstore <secretstore-name> -n <namespace>
==> kubectl delete secretstore vault-backend -n student-api-ns

Or if you are using clusterSecretsStore
Apply the ClusterSecret command
==> kubectl apply -f cluster-secret-store.yaml 
==> kubectl get clustersecretstore -n student-api-ns
==> kubectl describe clustersecretstore vault-cluster-store -n student-api-ns
==> kubectl delete clustersecretstore vault-cluster-store -n student-api-ns

Apply the ExternalSecretStore Configuration
In your ExternalSecret, ensure you reference the SecretStore and the correct namespaces.
==> kubectl apply -f external-secret.yaml
==> kubectl get externalsecret -n student-api-ns
==> kubectl describe externalsecret postgres-secrets -n student-api-ns
==> kubectl delete externalsecret postgres-secrets -n student-api-ns

Step 7: Deploy the database manifest that contain all neccesary resources:
==> kubectl apply -f db-deployment.yaml

Step 8: Deploy the API manifest that contain all neccessary resources:
==> kubectl apply -f student-api-deployment.yaml

check the init container:
==> kubectl logs <pod_name> -c init-db -n student-api-ns

Check all resource:
==> kubectl get deployment,svc,pod,secret -n student-api-ns

Check for the access link or url
==> minikube -p sre-project service student-api-service -n student-api-ns --url