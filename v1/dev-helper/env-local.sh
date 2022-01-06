#!/bin/sh

export GIN_MODE=debug
export LIBRARY_APP_MODE="develop"
export LIBRARY_URL_SERVER_STORE_MEDIAS="https://xxxx.com"

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

# auth tokens
export LIBRARY_ACCESS_TOKEN_SECRET_KEY="1029azyb8374cwdv"
export LIBRARY_ACCESS_TOKEN_AUDIENCE="https://*.library.com"
export LIBRARY_ACCESS_TOKEN_ISSUER="https://www.library.com"
export LIBRARY_ACCESS_TOKEN_LIVE="1420" # 7 minutos. Tiempo antes de expirar el token de acceso
export LIBRARY_REFRESH_TOKEN_SECRET_KEY="1029azyb8374cwdv"
export LIBRARY_REFRESH_TOKEN_LIVE="3196800" # 37 dias. Tiempo antes de expirar el token de refrescar el access token

# store uploaded files=> system-files local
export LIBRARY_STORE_UPLOADED_FILES_MODE="files-system-local" # values: files-system-local or aws-s3
export LIBRARY_DESTINATION_TARGET="../gitignore/store"

# # store uploaded files=> aws-s3
# export LIBRARY_STORE_UPLOADED_FILES_MODE="aws-s3" # values: files-system-local or aws-s3
# export LIBRARY_DESTINATION_TARGET="cdn.goollodging.com"