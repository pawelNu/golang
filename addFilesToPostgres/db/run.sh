#! /bin/bash

docker build docker image
docker build -t my_schema .

docker run docker container
docker run --name my_schema -p 5432:5432 -d my_schema
