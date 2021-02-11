package ProgrammerController

import (
	"encoding/json"
	"log"
	"net/http"
	"postgresql/db"
	"postgresql/models"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

var err interface{}

func Log(err error, msg interface{}) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func List(c *gin.Context) {

	db := db.RUN().MustBegin()

	programmers := []models.Programmer{}

	err = db.Select(&programmers, "SELECT * FROM programmer;")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "failed",
			"message": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "all list programmer",
		"data":   programmers,
	})
	//=======================
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	Log(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	Log(err, "Failed to open a channel")
	defer ch.Close()

	Log(err, "Failed to declare a queue")

	err = ch.ExchangeDeclare(
		"transaction", // name
		"topic",       // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	Log(err, "Failed to declare an exchange")
	//=======================

	jsonResponse, err := json.Marshal(&programmers)
	if err != nil {
		Log(err, "Failed to publish a message")
	}

	err = ch.Publish(
		"transaction", // exchange
		"list",        // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/aplain",
			Body:        []byte(jsonResponse),
		})

	Log(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", jsonResponse)
}

func Read(c *gin.Context) {

}

func EndTask(c *gin.Context) {

}
