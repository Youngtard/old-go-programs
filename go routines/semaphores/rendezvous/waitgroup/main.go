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

// Main Thread/Thread 1
func main() {
	go secondThread()
	firstSem.Add(1)
	// Wait 1
	firstSem.Wait()
	// Display 1
	fmt.Printf(firstOutput + "\n")

	scanner.Scan()
	// Obtain
	secondOutput = scanner.Text()
	secondSem.Done()

}

// Thread 2
func secondThread() {
	secondSem.Add(1)
	scanner.Scan()
	// Obtain 1
	firstOutput = scanner.Text()
	// Signal 1
	firstSem.Done()

	// Wait 2
	secondSem.Wait()
	// Display 2
	fmt.Printf(firstOutput + "\n")

}
