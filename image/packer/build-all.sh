#!/bin/bash

cd "$(git rev-parse --show-toplevel)"/image/packer || exit 1

packer build -var "esx_host=$GOVC_URL" -var "remote_username=$GOVC_USERNAME" -var "remote_password=$GOVC_PASSWORD" -on-error=abort image.json
