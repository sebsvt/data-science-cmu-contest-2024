package event

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/standardise-software/user-account-service/logs"
)

type accountEventHandler struct {
	channel *amqp.Channel
}

func NewAccountEventHandler(channel *amqp.Channel) EventHandler {
	return accountEventHandler{channel: channel}
}

// Handle implements EventHandler.
func (message accountEventHandler) Sender(name string, event_bytes []byte) error {
	q, err := message.channel.QueueDeclare(name, false, false, false, false, nil)
	if err != nil {
		logs.Error(err)
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := message.channel.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        event_bytes,
	}); err != nil {
		logs.Error(err)
		return err
	}
	logs.Info("Message has been sent to receiver")
	return nil
}
