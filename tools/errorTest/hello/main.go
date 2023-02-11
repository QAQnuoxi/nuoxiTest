package main

import (
	"app1/errorTest"
	"fmt"
	"log"
)

func hello() {
	log.SetPrefix("hello:")
	log.SetFlags(0)
	message, error1 := errorTest.Hello("tat")
	if error1 != nil {
		log.Fatal(error1)
	}
	fmt.Println(message)
}
