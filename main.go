package main

import (
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/AlpacaLabs/api-hermes/internal/app"
	"github.com/AlpacaLabs/api-hermes/internal/configuration"
)

func main() {
	c := configuration.LoadConfig()

	logrus.Infof("Loaded config: %s", c)

	a := app.NewApp(c)

	var wg sync.WaitGroup

	wg.Add(1)
	go a.Run()

	wg.Wait()
}
