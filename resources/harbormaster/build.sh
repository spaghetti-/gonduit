#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

rm -r /home/drydock/go/src

mkdir -p /home/drydock/go/src/github.com/spaghetti-
ln -sf $(pwd) /home/drydock/go/src/github.com/spaghetti-/
cd /home/drydock/go/src/github.com/spaghetti-/gonduit

glide install
go build $(glide novendor)
go test $(glide novendor)
