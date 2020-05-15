package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Topic      string   `default:"Topic1"`
	Partition  int      `default:1`
	BrokerURLs []string `default:"localhost:19092,localhost:29092,localhost:39092"`
	ClientID   string   `default:"kafka-client-id"`
	ListenAddr string   `default:"0.0.0.0:9000"`
}

type KafkaParty struct {
	conn *kafka.Conn
	cfg  Config
}

func produceMessages() error {
	topic := "Topic1"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "0.0.0.0:9092", topic, partition)
	if err != nil {
		return err
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	return nil
}

type KafkaMessage struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func ConnectKafka(cfg Config) (*kafka.Conn, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "0.0.0.0:9092", cfg.Topic, cfg.Partition)
	if err != nil {
		return conn, err
	}
	err = conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		return conn, err
	}

	return conn, nil
}

var writer *kafka.Writer

func ConfigureKafka(brokerUrls []string, clientID string, topic string) (w *kafka.Writer, err error) {
	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: clientID,
	}
	log.Info("Broker urls: ", brokerUrls)
	config := kafka.WriterConfig{
		Brokers:          brokerUrls,
		Topic:            topic,
		Balancer:         &kafka.LeastBytes{},
		Dialer:           dialer,
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}

	w = kafka.NewWriter(config)
	writer = w
	return w, nil
}

func PushKafka(ctx context.Context, key, value []byte) error {
	message := kafka.Message{
		Key:   key,
		Value: value,
	}

	return writer.WriteMessages(ctx, message)
}

func main() {
	cfg := Config{
		Topic:      "Topic1",
		Partition:  1,
		BrokerURLs: []string{"localhost:9092", "localhost:9093"},
		ClientID:   "kafka-client-id",
		ListenAddr: "0.0.0.0:9000",
	}
	var kp KafkaParty
	kp.cfg = cfg

	// configure kafka
	kafkaProducer, err := ConfigureKafka(cfg.BrokerURLs, cfg.ClientID, cfg.Topic)
	if err != nil {
		log.WithError(err).Error("could not create kafkaproducer")
	}
	defer kafkaProducer.Close()

	e := echo.New()
	e.POST("/", kp.ProduceMessage)
	e.Logger.Info("started prodyucer at 8080")
	e.Logger.Fatal(e.Start(":6789"))
}

func (kp KafkaParty) ProduceMessage(c echo.Context) error {
	km := new(KafkaMessage)
	err := c.Bind(km)
	if err != nil {
		log.WithError(err).Error("producemessage: could not bind body")
		return c.String(http.StatusBadRequest, "oh no, we require a body consisting of key and value")
	}
	kafkamessageInBytes, err := json.Marshal(km)
	if err != nil {
		log.WithError(err).Error("producemessage: could not marshal kafkamessage")
		return c.String(http.StatusInternalServerError, "not able to marshal kafka message")
	}
	err = PushKafka(c.Request().Context(), []byte("paap"), kafkamessageInBytes)
	if err != nil {
		log.WithError(err).Error("producemessage: could not push kafkamessage")
		return c.String(http.StatusInternalServerError, "not able to push to kafka")
	}
	return c.String(http.StatusOK, "ok")
}
