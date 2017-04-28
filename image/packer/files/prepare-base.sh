#!/bin/bash

brwhte="$(tput setaf 15)"
brblue="$(tput setaf 12)"
bryllw="$(tput setaf 11)"
brprpl="$(tput setaf 13)"
creset="$(tput sgr0)"

warrow="${brwhte}=>${creset}"
barrow="${brblue}==>${creset}"
yarrow="${bryllw}===>${creset}"

BASE_DIR=$(dirname $(readlink -f "$BASH_SOURCE"))

echo "${warrow} Updating packages"
tdnf distro-sync --refresh -y
echo "${warrow} Installing extra packages"
echo "${barrow} Getting tar linux-esx ethtool gawk socat"
tdnf install -y tar linux-esx ethtool gawk socat

# install bash-completion

echo "${warrow} Configuring OS"
echo "${barrow} Installing bash-completion"
tar -C /usr/share -xf /tmp/bash-completion.tar.xz
echo '
# Use bash-completion, if available
if [[ $PS1 && -f /usr/share/bash-completion/bash_completion ]]; then
  . /usr/share/bash-completion/bash_completion
fi
' > /etc/profile.d/bash-completions.sh

echo "${barrow} Installing jq"
curl -o /usr/bin/jq -L'#' https://github.com/stedolan/jq/releases/download/jq-1.5/jq-linux64
chmod +x /usr/bin/jq

echo "${barrow} Installing goodhosts"
latestgoodhosts=$(curl -s https://api.github.com/repos/casualjim/goodhosts/releases/latest | jq -r .tag_name)
curl -o /usr/bin/goodhosts -L'#' https://github.com/casualjim/goodhosts/releases/download/${latestgoodhosts}/goodhosts
chmod +x /usr/bin/goodhosts

echo "${barrow} Installing setup-network-environment"
latestnetenv=$(curl -s https://api.github.com/repos/kelseyhightower/setup-network-environment/releases/latest | jq -r .tag_name)
curl -o /usr/bin/setup-network-environment -L'#' https://github.com/kelseyhightower/setup-network-environment/releases/download/${latestnetenv}/setup-network-environment
chmod +x /usr/bin/setup-network-environment

echo "${barrow} Installing govc"
latestgovc=$(curl -s https://api.github.com/repos/vmware/govmomi/releases/latest | jq -r .tag_name)
curl -sSL'#' https://github.com/vmware/govmomi/releases/download/${latestgovc}/govc_linux_amd64.gz | gunzip -c - > /usr/bin/govc 
chmod +x /usr/bin/govc

echo "${barrow} Configuring grub"
sed -i 's/set timeout=5/set timeout=0/' /boot/grub2/grub.cfg
sed -i 's/insmod gfxterm/#insmod gfxterm/' /boot/grub2/grub.cfg
sed -i 's/insmod png/#insmod png/' /boot/grub2/grub.cfg
sed -i 's/set gfxmode/#set gfxmode/' /boot/grub2/grub.cfg
sed -i 's/gfxpayload=keep/#gfxpayload=keep/' /boot/grub2/grub.cfg
sed -i 's/terminal_output/#terminal_output/' /boot/grub2/grub.cfg
sed -i 's/set theme/#set theme/' /boot/grub2/grub.cfg
sed -i 's/quiet //' /boot/photon.cfg

echo "${barrow} Configuring password expiration"
chage -I -1 -m 0 -M 99999 -E -1 root

echo "${barrow} Add docker group"
groupadd docker
