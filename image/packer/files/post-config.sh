#!/bin/bash

warrow="=>"
barrow="==>"
yarrow="===>"

BASE_DIR=$(dirname $(readlink -f "$BASH_SOURCE"))

configure_haveged() {
  echo "==> Configuring haveged"
  systemctl daemon-reload
  systemctl enable haveged
}

configure_docker() {
  echo "==> Configuring docker"
  systemctl daemon-reload
  systemctl enable docker
}

configure_sshd() {
  echo "UseDNS no" >> /etc/ssh/sshd_config
}

configure_lvm() {
  echo "==> Configuring device mapper for docker"
  pvcreate /dev/sdb
  vgcreate docker /dev/sdb
  lvcreate --wipesignatures y -n thinpool docker -l 95%VG
  lvcreate --wipesignatures y -n thinpoolmeta docker -l 1%VG
  lvconvert -y --zero n -c 512K --thinpool docker/thinpool --poolmetadata docker/thinpoolmeta

  lvchange --metadataprofile docker-thinpool docker/thinpool

  lvchange --metadataprofile docker-thinpool docker/thinpool
}

configure_sshd
configure_haveged
configure_lvm
configure_docker
