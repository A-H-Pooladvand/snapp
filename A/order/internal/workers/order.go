package workers

import (
	"google.golang.org/protobuf/proto"
	"order/internal/models"
	"order/pkg/log"
	"order/pkg/rabbit"
	pb "order/proto"
)

func Run() {
	conn, err := rabbit.Dial()

	if err != nil {
		log.Error(err)
		panic(err)
	}

	ch, err := conn.Channel()

	if err != nil {
		log.Error(err)
		panic(err)
	}

	q, err := ch.QueueDeclare(
		"order", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	if err != nil {
		log.Error(err)
		panic(err)
	}

	// Consume messages
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		log.Error(err)
		panic(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			// Deserialize the message using protobuf
			message := new(pb.Order)
			err := proto.Unmarshal(d.Body, message)
			order := models.NewOrderFromProto(message)

			if order.Persist().IsEmpty() {
				log.Error("Unable to save order", message)
			}

			if err != nil {
				log.Error("Failed to deserialize message: %s", err)
				continue
			}

			log.Error("Received a message: %+v", message)
		}
	}()

	log.Error(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
