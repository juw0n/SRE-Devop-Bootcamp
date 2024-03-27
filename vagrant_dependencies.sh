#!/bin/bash

echo -e "\e[0;36mDocker and Docker Compose installation started!\e[0m"

echo -e "\e[0;36mRun apt-update first!\e[0m"
# Update package lists
sudo apt-get update

# For non-Gnome Desktop environments, gnome-terminal must be installed:
echo -e "\e[0;36mor non-Gnome Desktop environments, Intall gnome-terminal!\e[0m"
sudo apt install gnome-terminal

# Add Docker's official GPG key:
echo -e "\e[0;36mAdd Docker official GPG Key!\e[0m"
sudo apt-get update
sudo apt-get install -y ca-certificates curl
sudo install -y -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc

# Add the repository to Apt sources:
echo -e "\e[0;36mAdd the repository to thr Apt!\e[0m"
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# Update package lists again
echo -e "\e[0;36mUpdate packages again!\e[0m"
sudo apt-get update

# Install the Docker packages latest version
echo -e "\e[0;36mInstall the docker latest packages!\e[0m"
sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

echo -e "\e[0;36mYou have now successfully installed and started Docker Engine and Docker Compose!\e[0m"

echo -e "\e[0;36m-------------------------------------------------\e[0m"

echo -e "\e[0;36mInstalling make...\e[0m"

# Install the 'make' package
echo -e "\e[0;36mInstall make Package!\e[0m"
sudo apt-get install -y make

echo -e "\e[0;36mmake installation complete!\e[0m"

echo -e "\e[0;36m-------------------------------------------------\e[0m"
# Add User to Docker Group to grant permission to run docker command
echo -e "\e[0;36m----ssh into the VM----\e[0m"
echo -e "\e[0;36m----RUN 'sudo usermod -aG docker $USER'----\e[0m"
echo -e "\e[0;36m----Logout and Login again after the VM bootup----\e[0m"