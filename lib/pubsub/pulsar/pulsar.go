package pulsar

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
)

type Pulsar struct {
	producer pulsar.Producer
}

func (p *Pulsar) Publish(ctx context.Context, data []byte) error {
	_, err := p.producer.Send(ctx, &pulsar.ProducerMessage{
		Payload: data,
	})
	return err
}

func (p *Pulsar) Close() {
	p.producer.Close()
}

func New(url, topic string) (*Pulsar, error) {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: url,
	})
	if err != nil {
		return nil, err
	}

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
	})
	if err != nil {
		return nil, err
	}

	return &Pulsar{
		producer: producer,
	}, nil
}
