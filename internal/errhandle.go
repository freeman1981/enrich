package internal

import "log"

func errHandle(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
