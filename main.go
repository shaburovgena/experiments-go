package main

import (
	"experiments/internal/app/apiserver"
	"log"
)

func main() {
	config := apiserver.CreateConfig()

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
