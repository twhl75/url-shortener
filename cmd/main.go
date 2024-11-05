package main

import (
	"fmt"
	"os"

	"github.com/twhl75/url-shortener/internal/controller"
	"github.com/twhl75/url-shortener/internal/service"
)


func run() error {
	service := service.NewService()

	controller := controller.New(service)

	controller.Run()

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}



