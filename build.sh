#!/bin/bash

go build .
mv nmos-node-generator /var/www/nmos-node-generator/
env GOOS=windows GOARCH=amd64 go build .
mv nmos-node-generator.exe /var/www/nmos-node-generator/
cp README.md /var/www/nmos-node-generator/
cp config.json /var/www/nmos-node-generator/