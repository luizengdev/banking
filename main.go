package main

import (
	"github.com/luizengdev/banking/app"
	"github.com/luizengdev/banking/logger"
)

func main() {

	logger.Info("Starting the application...")
	app.Start()
}
