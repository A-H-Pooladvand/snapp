package order

import (
	"encoding/json"
	"gateway/pkg/log"
	"gateway/pkg/rabbit"
	"gateway/pkg/res"
	"gateway/pkg/util"
	"github.com/go-playground/validator/v10"
	"github.com/streadway/amqp"
	"net/http"
)

func (o Controller) Create(w http.ResponseWriter, r *http.Request) {
	request := new(Request)

	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		res.Response(w).BadRequest("Unable to process the body of request")
		return
	}

	// Just a minimal messy validation nothing fancy
	if err := validator.New().Struct(request); err != nil {
		res.Response(w).BadRequest(err.Error())
		return
	}

	conn, err := rabbit.Dial()

	if err != nil {
		log.Error(err)
		res.Response(w).ServerError("Something went wrong please contact support")
		return
	}

	defer util.Closer(conn)

	ch, err := conn.Channel()

	if err != nil {
		log.Error("Failed to open a channel: %v", err)
		res.Response(w).ServerError("Something went wrong please contact support")
		return
	}

	defer util.Closer(ch)

	q, err := ch.QueueDeclare(
		o.Queue(), // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	if err != nil {
		log.Error("Failed to declare a queue: %v", err)
		res.Response(w).ServerError("Something went wrong please contact support")
		return
	}

	msg, err := request.Serialize()

	if err != nil {
		log.Error("Unable to create proto message", err)
		res.Response(w).ServerError("Something went wrong please contact support")
		return
	}

	err = ch.Publish(
		o.Exchange(), // exchange
		q.Name,       // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/protobuf",
			Body:        msg,
		})

	if err != nil {
		log.Error("Failed to publish a message: %v", err)
		panic(err)
	}
}
