#!/bin/bash

# Add ansible user
useradd -m -s /bin/bash ansible
usermod -aG sudo ansible
mkdir -p /home/ansible/.ssh
echo ${SSH_KEY_PUBLIC_HOME_DESKTOP} >> /home/ansible/.ssh/authorized_keys
echo ${SSH_KEY_PUBLIC_GITHUB_ACTIONS} >> /home/ansible/.ssh/authorized_keys
chown -R ansible:ansible /home/ansible/.ssh
echo "ansible ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers

# Add user
useradd -m -s /bin/bash medvesekg
usermod -aG sudo medvesekg
mkdir -p /home/medvesekg/.ssh
echo ${SSH_KEY_PUBLIC_HOME_DESKTOP} >> /home/medvesekg/.ssh/authorized_keys
chown -R medvesekg:medvesekg /home/medvesekg/.ssh
echo "medvesekg ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers

# Disable SSH password authentiaction
sed -i -E 's/^#?PasswordAuthentication.*/PasswordAuthentication no/' /etc/ssh/sshd_config
rm /etc/ssh/sshd_config.d/50-cloud-init.conf # this file allows password auth and needs to be deleted

# Disable root login
sed -i -E 's/^#?PermitRootLogin.*/PermitRootLogin no/' /etc/ssh/sshd_config

service ssh restart # restart ssh for changes to take effect

apt-get update
apt-get install -y python3 python3-pip python3-venv
python3 -m venv /home/ansible/venv