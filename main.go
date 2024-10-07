package main

import (
	"fmt"
	"log"

	"github.com/centrex/webcore/core/env"
	"github.com/rochi88/goapi/bootstrap"
)

func main() {
	app := bootstrap.Application()

	// Start server
	err := app.Listen(fmt.Sprintf(":%s", env.GetEnv("PORT", "3000")))
	if err != nil {
		log.Fatal(err)
	}
}
