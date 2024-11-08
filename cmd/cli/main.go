package main

import (
	"flag"
	"helper-service-bus/config"
	"helper-service-bus/internal/readFileJson"
	"helper-service-bus/internal/serviceBus"
)

func main() {
	connectionString, dirName, err := config.ReadEnviroment(".env")
	if err != nil {
		panic(err)
	}

	senderServiceBus := serviceBus.CreateSenderServiceBusQueue(connectionString)

	var queueName string
	flag.StringVar(&queueName, "q", "", "Queue name")
	flag.Parse()

	if queueName == "" {
		panic("Queue name is required")
	}

	message, success := readFileJson.ReadFileJsonToGetMessageBody(queueName, dirName)

	if !success {
		panic("Message not found")
	}

	senderServiceBus(queueName, message)
}
