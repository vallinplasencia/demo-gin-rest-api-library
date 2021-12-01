#! /bin/bash

echo "Building executable"
mkdir -p ../gitignore/executables # ruta donde se crean los executables
mkdir -p ../gitignore/logs #ruta donde se ponen los archivos de logs

echo "Import enviroment local"
. ./env-local.sh
# echo "Import enviroment production"
# . ./env-prod.sh

go run ../cmd/server/main.go

# go build -o ../gitignore/executables/server ../cmd/main.go
# echo "Run server"
# ./../gitignore/executables/server
