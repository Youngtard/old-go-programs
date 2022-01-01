package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

var consoleReady sync.WaitGroup
var consoleoutput string

func main() {
	consoleReady.Add(1)
	go thread2()
	consoleReady.Wait()
	fmt.Println(consoleoutput)

}

func thread2() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	consoleoutput = scanner.Text()
	consoleReady.Done()
}
