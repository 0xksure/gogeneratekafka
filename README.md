# I am a kafka producer

My purpose is simple - produce data and store it in a kafka topic
...in go of course

## TODO:

- Connect a database to kafka
- create consumer and producer endpoints in go

## Local setup of zookeeper and kafka with docker compose

## Multibroker kafka

### Listen to topic

Download [Kafkacat](https://github.com/edenhill/kafkacat)

### Spin those containers

> docker-compose up --build

### Publish message to topic

Run the command in one terminal

> kafkacat -b localhost:9093,localhost:9093 -t Topic1 -P

and write whatever you like. Forexample

- Wow! Kafka is way more confusing than presented in the blogs
- Now that I have the solution lets search for the problem

### Consume message from topic

Run command in the other terminal

> kafkacat -b localhost:9093,localhost:9093 -t Topic1 -C
