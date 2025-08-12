package main

import (
	"aleksale/webchecker/checker"
)

// func fetchUrl(url string, done chan bool) {
// 	fmt.Printf("Fetching url %s\n", url)
// 	time.Sleep(2 * time.Second)
// 	fmt.Printf("Fetched url %s\n", url)

// 	done <- true
// }

// func main() {
// 	fmt.Println("Start fetching...")

// 	done := make(chan bool)

// 	go fetchUrl("Url1", done)
// 	go fetchUrl("Url2", done)
// 	go fetchUrl("Url3", done)

// 	for i := 0; i < 3; i++ {
// 		<-done
// 	}

// 	fmt.Println("Done fetchin!.")
// }

// ===================================================
// ===================================================

// func sender(ch chan string, done chan bool) {
// 	for i := 1; i <= 3; i++ {
// 		ch <- fmt.Sprintf("message %d", i)
// 		time.Sleep(100 * time.Millisecond)
// 	}
// 	close(ch) // Close the channel when done sending
// 	done <- true
// }

// func receiver(ch chan string, done chan bool) {
// 	// runs until the channel is closed
// 	for msg := range ch {
// 		fmt.Println("Received:", msg)
// 	}
// 	done <- true
// }

// func main() {
// 	ch := make(chan string)
// 	senderDone, receiverDone := make(chan bool), make(chan bool)

// 	go sender(ch, senderDone)
// 	go receiver(ch, receiverDone)

// 	// Wait for both sender and receiver to complete
// 	<-senderDone
// 	<-receiverDone

// 	fmt.Println("All operations completed!")
// }

func main() {
	checker.WebStatusCheck()
}
