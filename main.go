package main

import (
	"flag"
	"log"
)

func main() {
	mode := flag.String("mode", "sub", "Mode to run: 'sub' (subscriber) or 'pub' (publisher)")
	flag.Parse()

	switch *mode {
	case "sub":
		log.Println("Starting subscriber...")
		Subscriber()
	case "pub":
		log.Println("Publishing message...")
		Publish()
	default:
		log.Fatalf("unknown mode: %s (use 'sub' or 'pub')", *mode)
	}
}
