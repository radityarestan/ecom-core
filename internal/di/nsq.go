package di

import (
	"github.com/nsqio/go-nsq"
	"github.com/radityarestan/ecom-authentication/internal/pkg"
	"github.com/radityarestan/ecom-authentication/internal/shared/config"
)

func NewNSQProducer(conf *config.Config) (np *pkg.NSQProducer, err error) {
	np = &pkg.NSQProducer{}
	np.Env = conf.NSQ

	nsqConfig := nsq.NewConfig()
	np.Producer, err = nsq.NewProducer(np.Env.Host+":"+np.Env.Port, nsqConfig)
	if err != nil {
		return nil, err
	}

	return np, nil
}
