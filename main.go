package main

import (
	"fmt"
	"log"

	"github.com/rochi88/goapi/app/helpers"
	"github.com/rochi88/goapi/bootstrap"
)

func main() {
	app := bootstrap.Application()

	// Start server
	err := app.Listen(fmt.Sprintf(":%s", helpers.GetEnv("PORT", "3000")))
	if err != nil {
		log.Fatal(err)
	}
}
