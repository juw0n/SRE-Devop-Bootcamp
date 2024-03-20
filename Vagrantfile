Vagrant.configure("2") do |config|
    config.vm.box = "ubuntu/jammy64"
    config.vm.provision "shell", inline: <<-SHELL
        bash /vagrant/vagrant_dependencies.sh
    SHELL
    config.vm.network :forwarded_port, guest: 80, host: 80
end