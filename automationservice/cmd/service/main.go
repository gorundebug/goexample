package main

import (
	"context"
	"flag"
	"log"
	"os"

	automationservice "github.com/gorundebug/automationservice/internal/app"
)

func main() {
	valuesPathArg := flag.String("values", "./config/overrides.yaml", "service config values path")
	configPathArg := flag.String("config", "./config/config.yaml", "service config path")
	flag.Parse()

	if err := automationservice.Start(
		context.Background(),
		nil,
		nil,
		configPathArg,
		valuesPathArg,
		os.Getenv("APP_VERSION"),
		os.Getenv("APP_COMMIT"),
	); err != nil {
		log.Printf("start service error: %s", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
