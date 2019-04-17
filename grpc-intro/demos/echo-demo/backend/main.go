package main

import (
	"fmt"
	"time"

	_ "google.golang.org/grpc"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(30 * time.Second)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
