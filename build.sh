#!/bin/bash

go build .
cp nmos-node-generator /mnt/c/Users/irowell/Documents/nmos_node_generator/
cp config.json /mnt/c/Users/irowell/Documents/nmos_node_generator/
env GOOS=windows GOARCH=amd64 go build .
