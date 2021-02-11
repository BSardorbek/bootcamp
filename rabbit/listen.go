package main

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func Log(err error, msg interface{}) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//Receiver ...
func main() {

	// conn, err := amqp.Dial("")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	Log(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	Log(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"transaction", // name
		"topic",       // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)

	Log(err, "Failed to ExchangeDeclare ")

	q, err := ch.QueueDeclare(
		"test", // name
		true,   // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)

	Log(err, "Failed to declare a queue")

	routingkey := os.Args[1]

	err = ch.QueueBind(
		q.Name,        // queue name
		routingkey,    // routing key
		"transaction", // exchange
		false,
		nil,
	)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	Log(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {

			// _, ok := d.Headers["x-key-name"]

			// if !ok {
			// 	log.Printf("failed a message: %s", d.Body)
			// 	d.Nack(false, false)
			// }
			log.Printf("Received a message: %s", d.Body)
			d.Nack(false, false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
