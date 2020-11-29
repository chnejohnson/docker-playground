### Kafka

- run ZooKeeper

`docker run --name zookeeper --rm -p 2181:2181 -d zookeeper`

- build images of mykafka

`docker build -t mykafka .`

- run images on port 9092

`docker run -p 9092:9092 --rm mykafka`