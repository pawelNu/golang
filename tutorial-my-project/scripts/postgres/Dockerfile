# Dockerfile

# Use the official PostgreSQL image with the 16.1 version
FROM postgres:16.1

# Define an environment variable with the password for the "postgres" user
ENV POSTGRES_PASSWORD=test

# Configure ports
EXPOSE 5432

# Copy the SQL initialization file to the /docker-entrypoint-initdb.d/ directory
COPY create_database.sql /docker-entrypoint-initdb.d/
