#!/bin/bash

DBNAME="dousheng"  #数据库名称

#创建数据库

drop_db_sql="drop database if EXISTS ${DBNAME}"
create_db_sql="create database IF NOT EXISTS ${DBNAME}"

mysql -h ${MYSQL_HOST} -P ${MYSQL_PORT} -u ${MYSQL_USER} -p${MYSQL_PASSWORD} -e "${drop_db_sql}"
mysql -h ${MYSQL_HOST} -P ${MYSQL_PORT} -u ${MYSQL_USER} -p${MYSQL_PASSWORD} -e "${create_db_sql}"

echo "exit"