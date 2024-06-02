FROM mcr.microsoft.com/devcontainers/go:0-1.16

# Install PostgreSQL and Redis
RUN apt-get update && \
    apt-get install -y postgresql postgresql-contrib redis-server && \
    apt-get clean

# Add setup script
COPY setup.sh /usr/local/bin/setup.sh
RUN chmod +x /usr/local/bin/setup.sh

# Expose necessary ports
EXPOSE 5432 6379 8080
