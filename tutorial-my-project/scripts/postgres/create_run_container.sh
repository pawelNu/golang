dbName="rssagg_db"

# build docker image
docker build -t "$dbName" .

# run docker container
docker run --name "$dbName" -p 5431:5432 -d "$dbName"
