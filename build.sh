#!/bin/bash

go build .
cp nmos-node-generator /mnt/c/Users/irowell/Documents/nmos_node_generator/
# cp config.json /mnt/c/Users/irowell/Documents/nmos_node_generator/
# mv nmos-node-generator /var/www/nmos-node-generator/
env GOOS=windows GOARCH=amd64 go build .
# mv nmos-node-generator.exe /var/www/nmos-node-generator/
# cp README.md /var/www/nmos-node-generator/
# cp config.json /var/www/nmos-node-generator/