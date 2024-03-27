Vagrant.configure("2") do |config|
    config.vm.box = "ubuntu/jammy64"
    # Provisioning with shell script
    config.vm.provision "shell" do |shell|
        shell.path = "./vagrant_dependencies.sh"
    end
    # Forwarding port 80
    # I had to change the host port to 8084 for vagrant. 
    config.vm.network :forwarded_port, guest: 8080, host: 8084
end