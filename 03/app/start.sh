#!/bin/sh

export HW03_DB_USER=hw03
export HW03_DB_PASSWD=hw03
export HW03_DB_HOST=localhost
export HW03_DB_NAME=hw03
export HW03_LISTEN_PORT=8888
export EMAIL_CHECKER_HOST=localhost:8800
P=$(pwd)/$(dirname $0)
$P/arch-hw03