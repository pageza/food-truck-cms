#!/bin/bash

# Install PostgreSQL
sudo apt-get update
sudo apt-get install -y postgresql postgresql-contrib

# Start PostgreSQL service
sudo service postgresql start

# Install Redis
sudo apt-get install -y redis-server

# Start Redis service
sudo service redis-server start

# Create PostgreSQL user and database
sudo -u postgres psql -c "CREATE USER devuser WITH PASSWORD 'devpassword';"
sudo -u postgres psql -c "ALTER USER devuser WITH SUPERUSER;"
sudo -u postgres psql -c "CREATE DATABASE devdb OWNER devuser;"

# Allow remote connections to PostgreSQL
echo "host all  all    0.0.0.0/0  md5" | sudo tee -a /etc/postgresql/12/main/pg_hba.conf
echo "listen_addresses='*'" | sudo tee -a /etc/postgresql/12/main/postgresql.conf
sudo service postgresql restart
