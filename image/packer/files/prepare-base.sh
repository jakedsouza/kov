#!/bin/bash

warrow="=>"
barrow="==>"
yarrow="===>"

BASE_DIR=$(dirname $(readlink -f "$BASH_SOURCE"))

configure_os() {

  echo "${warrow} Updating packages"
  tdnf distro-sync --refresh -y
  echo "${warrow} Installing extra packages"
  tdnf install -y tar \
    haveged \
    linux-esx \
    ethtool \
    gawk \
    socat \
    lvm2 \
    thin-provisioning-tools

  echo "${warrow} Configuring OS"
  echo "${barrow} Installing bash-completion"
  tar -C /usr/share -xf /tmp/bash-completion.tar.xz

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
}

install_etcd() {
  ETCD_VER=$(curl -s https://api.github.com/repos/coreos/etcd/releases/latest | jq -r .tag_name)
  echo "installing etcd ${ETCD_VER}"
  echo "downloading etcd ${ETCD_VER}"
  DOWNLOAD_URL=https://github.com/coreos/etcd/releases/download
  curl -L'#' ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -o /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
  mkdir -p /tmp/test-etcd && tar xzf /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz -C /tmp/test-etcd --strip-components=1

  echo "installing etcd binaries ${ETCD_VER}"
  mv /tmp/test-etcd/{etcd,etcdctl} /usr/bin
  mkdir -p /var/lib/etcd
  rm -rf /tmp/test-etcd
}

install_kubernetes() {
  ver=$(curl -sL https://storage.googleapis.com/kubernetes-release/release/stable.txt)
  echo "installing kubernetes $ver"

  for fl in kube-apiserver kube-controller-manager kube-scheduler kubectl kubelet kube-proxy; do
    echo "installing ${fl}"
    curl -o /usr/bin/$fl -L'#' https://storage.googleapis.com/kubernetes-release/release/$ver/bin/linux/amd64/$fl
    chmod +x /usr/bin/$fl
  done
}

configure_lvm() {
  echo "==> Configuring device mapper for docker"
  systemctl enable lvm2-monitor
  systemctl start lvm2-monitor
}


configure_os

install_etcd
install_kubernetes

configure_lvm

systemctl disable docker
systemctl reboot
