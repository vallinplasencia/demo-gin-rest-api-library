#!/bin/sh

export GIN_MODE=debug
export LIBRARY_APP_MODE="develop"

# file logs
# export LIBRARY_PATH_FILE_LOGS="../gitignore/logs/all.log"

# http and https
export LIBRARY_ADDRESS_HTTP=":8080"
export LIBRARY_ADDRESS_HTTPS=":4443"
export LIBRARY_PATH_CERT_HTTPS="../gitignore/certs/cert.pem"
export LIBRARY_PATH_KEY_HTTPS="../gitignore/certs/key.pem"
export LIBRARY_READ_TIMEOUT="11"
export LIBRARY_WRITE_TIMEOUT="11"

# db
export LIBRARY_DB_MYSQL_USER="my-user"
export LIBRARY_DB_MYSQL_PASSWORD="my-pass"
export LIBRARY_DB_MYSQL_ADDRESS="localhost:3306"
export LIBRARY_DB_MYSQL_DBNAME="library"
