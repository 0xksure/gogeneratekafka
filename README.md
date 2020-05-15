# I am a kafka producer

My purpose is simple - produce data and store it in a kafka topic
...in go of course

## Set up zookeeper and kafka with docker compose

## Multibroker kafka

### Listen to topic

Download Kafkacat

> kafkacat -C -b localhost:9093,localhost:9092 -t Topic1 -p 0

### Publish message to partition

> echo 'publish to partition 0' | kafkacat -P -b localhost:9093,localhost:9092 -t Topic1 -p 0
