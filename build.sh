#!/bin/bash

go build .
mv nmos-node-generator /var/www/nmos-node-generator/
cp README.md /var/www/nmos-node-generator/
cp config.json /var/www/nmos-node-generator/