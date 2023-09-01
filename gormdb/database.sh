#!/bin/bash

DATABASE_NAME="dousheng"

CREATE="create database ${DATABASE_NAME}"

mysql -h ${MYSQL_HOST} -P ${MYSQL_PORT} -u ${MYSQL_USER} -p${MYSQL_PASSWORD} -e "${CREATE}"

echo "exit"