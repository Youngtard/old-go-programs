package main

import (
	"fmt"
)

func main() {
	// See channels as a means of signalling
	// Default channels, without capacity, are unbuffered: sending something blocks the sender until the receiver reads from the channel
	// A receive from an unbuffered channel happens before the send on that channel completes.
	// Just like communication (send operation) is not complete until feedback is received (read operation)
	// So in short, writes to unbuffered channel happens only when there is some routine ***waiting*** to read from channel, else the write operation is blocked forever and leads to deadlock.
	lock := make(chan bool)

	// for i := 1; i < 7; i++ {
	// 	go worker(i, lock)

	// }

	// fmt.Println(<-lock) // At this point, the read operation blocks and waits for a send from a go routine, but there is no go routine running at this point. So deadlock! - all (none actually) goroutines are asleep/finished running

	go func() {
		lock <- true         // the sender blocks until the receiver has received the value (unbuffered channel)
		fmt.Println("show?") // It shows sometimes. It doesnt show because the above unblocks when its value is received, after which the main() func exits the program, so there may be no time to run println("show?")
	}()

	fmt.Println(<-lock) // But here, a 	go routine is running/has run and performed a send operation on the channel in question so it doesn't block.

	// time.Sleep(2 * time.Second)
	// time.Sleep(10 * time.Second)

}

// func worker(id int, lock chan bool) {
// 	fmt.Printf("%d wants the lock\n", id)
// 	lock <- true
// 	fmt.Printf("%d has the lock\n", id)
// 	time.Sleep(500 * time.Millisecond)
// 	fmt.Printf("%d is releasing the lock\n", id)
// 	lock <- true
// }
