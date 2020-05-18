#!/bin/bash

# JDBC Drivers
# ------------
# MySQL
cd /usr/share/java/kafka-connect-jdbc/
curl https://cdn.mysql.com/Downloads/Connector-J/mysql-connector-java-8.0.16.tar.gz | tar xz 
# MS SQL
cd /usr/share/java/kafka-connect-jdbc/
curl http://central.maven.org/maven2/com/microsoft/sqlserver/mssql-jdbc/7.0.0.jre8/mssql-jdbc-7.0.0.jre8.jar --output mssql-jdbc-7.0.0.jre8.jar
# # Oracle
cp /db-leach/jdbc/lib/ojdbc8.jar /usr/share/java/kafka-connect-jdbc
# Now launch Kafka Connect
sleep infinity &
/etc/confluent/docker/run 