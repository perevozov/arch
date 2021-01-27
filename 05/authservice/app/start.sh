#!/bin/sh

export DB_USER=hw05
export DB_PASSWD=hw05
export DB_HOST=localhost
export DB_NAME=hw05
export LISTEN_PORT=8888
P=$(pwd)/$(dirname $0)
$P/authservice