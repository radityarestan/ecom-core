package pkg

import (
	"github.com/nsqio/go-nsq"
	"github.com/radityarestan/ecom-authentication/internal/shared/config"
)

type NSQProducer struct {
	Producer *nsq.Producer
	Env      config.NSQConfig
}

func (np *NSQProducer) Publish(message []byte) error {
	return np.Producer.Publish(np.Env.Topic, message)
}

func (np *NSQProducer) Stop() {
	np.Producer.Stop()
}
