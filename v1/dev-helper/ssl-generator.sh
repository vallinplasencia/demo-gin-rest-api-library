#! /bin/bash

echo "Generating certs for https server"

mkdir -p ../gitignore/certs #ruta donde se ponen cert and key

openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ../gitignore/certs/key.pem -out ../gitignore/certs/cert.pem

echo "Generated certs"