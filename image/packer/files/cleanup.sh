#!/bin/bash -eux

SSH_USER=${SSH_USERNAME:-vagrant}

DISK_USAGE_BEFORE_CLEANUP=$(df -h)

echo "==> Cleaning up tmp"
rm -rf /tmp/*

# Cleanup tdnf cache
tdnf clean all

echo "==> Removing man pages"
rm -rf /usr/share/man/*
echo "==> Removing any docs"
rm -rf /usr/share/doc/*
echo "==> Removing caches"
find /var/cache -type f -exec rm -rf {} \;

# Remove Bash history
unset HISTFILE
rm -f /root/.bash_history
[ -f /home/vagrant/.bash_history ] && rm -f /home/${SSH_USER}/.bash_history

# Clean up log files
find /var/log -type f | while read f; do echo -ne '' > $f; done;

echo "==> Clearing last login information"
echo -ne '' >/var/log/lastlog
echo -ne '' >/var/log/wtmp
echo -ne '' >/var/log/btmp

echo -ne '' > /root/.bashrc

# Clear machine id
echo -ne '' > /etc/machine-id

# Zero out the free space to save space in the final image
dd if=/dev/zero of=/EMPTY bs=1M  || echo "dd exit code $? is suppressed"
rm -f /EMPTY

# Make sure we wait until all the data is written to disk, otherwise
# Packer might quit too early before the large files are deleted
sync

echo "==> Disk usage before cleanup"
echo ${DISK_USAGE_BEFORE_CLEANUP}

echo "==> Disk usage after cleanup"
df -h
