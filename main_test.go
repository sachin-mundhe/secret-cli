package main

import (
	"log"
	"testing"
)

func TestMain(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			log.Fatalln("Error occured in main function")
		}
	}()
	main()
}
