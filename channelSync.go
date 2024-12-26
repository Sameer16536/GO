package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second) // Simulate kaam hone mein lagne wala time
	fmt.Println("done")

	done <- true // Signal bhejna ki kaam complete ho gaya
}

func main() {
	done := make(chan bool, 1) // Ek buffered channel banaya

	go worker(done) // Worker function ek goroutine mein chal raha hai

	<-done // Yaha main function wait karega jab tak worker signal na bheje
}
