//What are Goroutines
// Lightweight Threads:
// ----Goroutines system threads ki tarah heavy nahi hote.
// ----Yeh Go runtime ke dwara manage hote hain.
//Concurrent Execution:

// Multiple kaam ek saath kar sakte ho.
// Jaise ek file read karte waqt dusra function database update kar sakta hai.
// go Keyword:

// Kisi bhi function ko goroutine mein execute karna ho, bas uske aage go lagao.

package main

import (
	"fmt"
	"time"
)

func bol(message string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(message, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	bol("Main Function")        // will run sequentially
	go bol("Goroutine A")       // will run concurrently
	go bol("Goroutine B")       // will run concurrently
	time.Sleep(3 * time.Second) // Main thread wait karega
	fmt.Println("Main Function Finished!")

}
