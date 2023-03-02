package utils

import (
	"log"
)

func ErrorWithPanic(err error) {
	defer recoverPanic()
	if err != nil {
		panic(err)
	}
}

func recoverPanic() {
	err := recover()
	if err != nil {
		log.Printf("| msg : %s", err)
	}
}
