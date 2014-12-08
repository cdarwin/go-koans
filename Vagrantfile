# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

PROVISION_SCRIPT = <<SCRIPT
  pacman -Sy --noconfirm go
SCRIPT

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "terrywang/archlinux"
  config.vm.provision "shell", inline: PROVISION_SCRIPT
end
