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
	firstSem.Add(1)

	ch := make(chan string)

	go firstThread(ch)
	go secondThread(ch)

	firstSem.Wait()

}

// Thread 1
func firstThread(ch chan string) {
	scanner.Scan()
	firstOutput = scanner.Text()
	ch <- firstOutput
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

	fmt.Printf(<-ch + "\n")
	firstSem.Done()

}
