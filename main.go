package main

import (
	"fmt"
	"github.com/radityarestan/ecom-authentication/internal/controller"
	"github.com/radityarestan/ecom-authentication/internal/di"
	"github.com/radityarestan/ecom-authentication/internal/shared"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var container = di.Container

	err := container.Invoke(func(deps shared.Deps, ch controller.Holder) error {
		var sig = make(chan os.Signal, 1)

		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

		ch.RegisterRoutes()

		go func() {
			deps.Logger.Infof("Starting server on port %s", deps.Config.Server.Port)
			if err := deps.Server.Start(fmt.Sprintf(":%s", deps.Config.Server.Port)); err != nil {
				deps.Logger.Errorf("Failed to start server: %v", err)
				sig <- syscall.SIGTERM
			}
		}()

		<-sig
		deps.Logger.Info("Shutting down server")
		deps.Close()
		return nil
	})

	if err != nil {
		panic(err)
	}
}
