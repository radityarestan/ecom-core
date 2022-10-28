package service

import "go.uber.org/dig"

type Holder struct {
	dig.In
	Auth    Auth
	Product Product
}

func Register(container *dig.Container) error {
	if err := container.Provide(NewAuth); err != nil {
		return err
	}

	if err := container.Provide(NewProduct); err != nil {
		return err
	}

	return nil
}
