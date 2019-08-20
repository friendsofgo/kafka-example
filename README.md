# Apache Kafka Example on Go

This is a simple repo to use [Apache Kafka](https://kafka.apache.org/) with Go, the code demo is based on a chat via
console line.

The code is divided into two part a client and a server.

## Run Apache Kafka

The `docker-compose.yml` is based on [gist by Matt Howlett](https://gist.githubusercontent.com/mhowlett/14f70af1a5b44fba80c9d1857a44bb98/raw/c0e6b15cf63801037cde9eb8987dbf767f7d06a2/docker-compose.yml) 

To run:

```sh
$ MY_IP=your-ip docker-compose up
```

If we want to create a topic, like `fogo-chat` you must run:

```sh
$ docker run --net=host --rm confluentinc/cp-kafka:5.0.0 kafka-topics --create --topic fogo-chat --partitions 4 --replication-factor 2 --if-not-exists --zookeeper localhost:32181
```

This command creates a topic named `fogo-chat` with `4 partitions` and `replication factor of 2`.

## Run the server

The default host is `http://localhost:8080`, you can change this configuration on `.env` file

```sh
make server-run
```

## Run the clients

You can run as many clients as you want, executing this:

```sh
make client-run
```