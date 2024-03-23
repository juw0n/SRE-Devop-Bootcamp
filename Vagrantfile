Vagrant.configure("2") do |config|
    config.vm.box = "ubuntu/jammy64"
    # Provisioning with shell script
    config.vm.provision "shell" do |shell|
        shell.path = "./vagrant_dependencies.sh"
    end
    # Forwarding port 80
    config.vm.network :forwarded_port, guest: 80, host: 80
end