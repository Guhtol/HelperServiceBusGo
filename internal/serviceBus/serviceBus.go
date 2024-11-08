package serviceBus

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

func CreateSenderServiceBusQueue(connectionString string) func(queueName, message string) {
	options := azservicebus.ClientOptions{}

	client, err := azservicebus.NewClientFromConnectionString(connectionString, &options)
	if err != nil {
		panic(err)
	}

	contextBackGround := context.Background()
	return func(queueName, message string) {
		sender, err := client.NewSender(queueName, nil)
		if err != nil {
			panic(err)
		}

		defer sender.Close(contextBackGround)

		sbMessage := azservicebus.Message{
			Body: []byte(message),
		}

		err = sender.SendMessage(context.TODO(), &sbMessage, nil)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Sent message: %s\n", message)
	}
}
