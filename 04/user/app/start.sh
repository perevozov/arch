#!/bin/sh

export DB_USER=hw03
export DB_PASSWD=hw03
export DB_HOST=localhost
export DB_NAME=hw03
export LISTEN_PORT=8888
export EMAIL_CHECKER_HOST=localhost:8800
P=$(pwd)/$(dirname $0)
$P/arch-userservice