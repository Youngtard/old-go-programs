package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

var firstOutput string
var secondOutput string

var firstSem sync.WaitGroup
var secondSem sync.WaitGroup

var scanner = bufio.NewScanner(os.Stdin)

// Main Thread
func main() {
	firstSem.Add(2)

	ch := make(chan string, 2)

	go firstThread(ch)
	go secondThread(ch)

	firstSem.Wait()
	firstSem.Wait()

	// select is used on multiple channels

	for i := range ch {
		fmt.Println(i)
	}
	// time.Sleep(time.Duration(5) * time.Second)
	// fmt.Println(len(ch))

	// for {
	// 	select {
	// 	case msg1 := <-ch:
	// 		fmt.Println(msg1)

	// 	case msg2 := <-ch:
	// 		fmt.Println(msg2)
	// 	}

	// }

}

// Thread 1
func firstThread(ch chan string) {
	scanner.Scan()
	secondOutput = scanner.Text()
	ch <- secondOutput
	firstSem.Done()
}

// Thread 2
func secondThread(ch chan string) {

	//To know if channel is closed
	// _, ok := <-ch
	// if ok {
	// 	// fmt.Printf(msg + "\n")
	// 	fmt.Printf(firstOutput + "\n")
	// 	firstSem.Done()
	// }
	scanner.Scan()
	firstOutput = scanner.Text()
	ch <- firstOutput
	firstSem.Done()

}
