package main

import "log"

func recoverable() {
	if x := recover(); x != nil {
		log.Println("panic:", x)
	}
}
