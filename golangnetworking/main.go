package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		//Program terminates with 1 as the code. If not included, the program will continue execution when we actually don't need more than 2 CL arguments
		os.Exit(1)
	}

	name := os.Args[1]

	addr := net.ParseIP(name)

	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Printf("The raw address is %s\n", name)
		fmt.Printf("The string address is %s\n", addr.String())

		// fmt.Println("The address is", addr.String(), name)
	}

	os.Exit(0)
}
