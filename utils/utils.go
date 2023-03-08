package utils

import "log"

func HandelERror(err error) {
	if err != nil {
		log.Panic(err)
	}
}
