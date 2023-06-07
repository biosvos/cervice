package main

import (
	"cervice/handler"
	"go-micro.dev/v4"
	"log"
)

func main() {
	service := micro.NewService(
		micro.Name("cervice"),
		micro.Handle(&handler.Handler{}),
	)
	service.Init()
	err := service.Run()
	if err != nil {
		log.Fatal(err)
	}
}
