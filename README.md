# gosampleserver

## What's this?

This is a sample app to show illustration for multi Docker Containers.

## Usage

### Start

```sh
docker build -t gosampleserver .
docker network create my-network
docker run -d --rm --name web -p 8080:8080 --network my-network gosampleserver
docker run -d --rm --name cache --network my-network memcached
# open http://0:8080 in your browser
```

### Stop

```sh
docker kill web
docker kill cache
docker network rm my-network
```

### with `docker-compose`

```sh
# start
docker-compose up -d
# stop
docker-compose down
```
