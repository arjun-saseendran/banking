package main

import (
	"github.com/arjun-saseendran/banking/app"
	"github.com/arjun-saseendran/banking/logger"
)

func main() {
	// log.Println("Starting application...")
	logger.Info("Starting application...")
	app.Start()
}
