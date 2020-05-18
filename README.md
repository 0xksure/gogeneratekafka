# I am a kafka producer

My purpose is simple - produce data and store it in a kafka topic
...in go of course

## Set up zookeeper and kafka with docker compose

## Multibroker kafka

### Listen to topic

Download Kafkacat

### Publish message to topic

> kafkacat -b localhost:9093,localhost:9093 -t Topic1 -P

### Consume message from topic

> kafkacat -b localhost:9093,localhost:9093 -t Topic1 -C
