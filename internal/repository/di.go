package repository

import "go.uber.org/dig"

type Holder struct {
	dig.In
	Auth Auth
}

func Register(container *dig.Container) error {
	if err := container.Provide(NewAuth); err != nil {
		return err
	}

	return nil
}
