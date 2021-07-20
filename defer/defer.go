package main

import (
	"errors"
	"log"
)

func main() {
	wasFail := false

	defer func() {
		if wasFail {
			log.Println("unsuccess")
			// metrics.WasUnssec
		} else {
			log.Println("success")
			// metrics.Success
		}
	}()

	log.Println("start")

	err := foo()
	if err != nil {
		wasFail = true
		log.Fatal("smthing wrong")
	}

	log.Println("finished")
}

func foo() error {
	return errors.New("some error")
}
