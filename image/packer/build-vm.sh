#!/bin/bash

vm_name=${1-base}
if govc vm.info -json=true $vm_name | jq -Me '.VirtualMachines[0].Name' 2>&1 >/dev/null; then
  echo "found the vm, destroying..."
  govc vm.destroy $vm_name
  govc datastore.rm $vm_name
fi

cd `git rev-parse --show-toplevel`/image/packer
(packer build -only=$vm_name -var "esx_host=$GOVC_URL" -var "remote_username=$GOVC_USERNAME" -var "remote_password=$GOVC_PASSWORD" -on-error=abort image.json)
govc import.ovf --pool '' ./$vm_name/$vm_name/$vm_name.ovf
