#!/bin/bash

echo "Docker and Docker Compose installation started!"

# Update package lists
sudo apt-get update

# Install prerequisites
apt-get install -y \
  apt-transport-https \
  ca-certificates \
  curl \
  software-properties \
  gnupg2

# Add Docker's official GPG key:
sudo apt-get install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc

# Add the repository to Apt sources:
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update

#Install the Docker packages.
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

# Install Docker Compose
sudo apt-get update
sudo apt-get install docker-compose-plugin


# Update package lists again
apt-get update

echo "You have now successfully installed and started Docker Engine and Docker Compose!"