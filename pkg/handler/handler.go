package handler

import (
	"context"
	"eff_mob_test/pkg/service"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	kafkaReader *kafka.Reader
	kafkaWriter *kafka.Writer
	services    *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	graphqlR := router.Group("/graphql")
	{
		graphqlR.GET("", h.PlaygroundHandler())
	}
	router.Any("/query", h.GraphqlHandler())

	api := router.Group("/user")
	{
		api.GET("", h.GetUser)
		api.POST("", h.CreateUser)
		api.DELETE("", h.DeleteUser)
		api.PUT("", h.UpdateUser)
	}

	return router
}

func (h *Handler) InitKafka(ctx context.Context) {
	r := h.InitKafkaReader()
	w := h.InitKafkaWriter()

	h.KafkaConsumeJSON(ctx, r, w)
}

func (h *Handler) InitKafkaReader() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "FIO",
		Partition: 0,
		MaxBytes:  10e3,
	})
}

func (h *Handler) InitKafkaWriter() *kafka.Writer {
	return &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "FIO_FAILED",
	}
}

func (h *Handler) KafkaConsumeJSON(ctx context.Context, r *kafka.Reader, w *kafka.Writer) {
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			logrus.Fatal(err)
		}
		if !json.Valid(msg.Value) {
			logrus.Errorln("error: JSON request format is invalid! Message is", string(msg.Value))
			if err := h.KafkaProduce(ctx, w, msg); err != nil {
				logrus.Errorln("Error while producing message to FIO_FAILED: ", err)
			}
			continue
		}
		err = h.services.CreateUser(msg.Value)
		if err != nil {
			logrus.Errorln(err)

			err = h.KafkaProduce(ctx, w, msg)

			if err != nil {
				logrus.Errorln(err)
			}
		}
	}
}

func (h *Handler) KafkaProduce(ctx context.Context, w *kafka.Writer, msg kafka.Message) error {
	err := w.WriteMessages(ctx, kafka.Message{
		Key:   msg.Key,
		Value: msg.Value,
	})
	if err != nil {
		return errors.New(fmt.Sprintln("could not write message: ", err.Error()))
	}
	logrus.Infoln("Message", string(msg.Value), "is written to FIO_FAILED")
	return nil
}
