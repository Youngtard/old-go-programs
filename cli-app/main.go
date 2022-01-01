package main

// CL Apps - Flags/Options, and CL Arguments

// go run prog -b -name Femi -age 21 <<<args>>>
// -b true(boy == true, b == boolean == true)

// We have flag arguments and non-flag arguments, positional arguments (https://gobyexample.com/command-line-flags)
// Write a more useful program implementing both types of args in its end-usage for better understanding by readers of article
// e.g. a git (cli?) clone - gitty or some other popular command line. google keep

// command, option, flag (booleantype option), argument. arguments tells the function what to act on/from??

// Read docs for solid info
// Confirm assertions/assumptions/statements

import (
	"flag"
	"fmt"
	"os"
)

// var name = flag.String("name", "World", "A name to say hello to.")
var name string

var spanish bool
var num int

// init all in main()
func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language.")
	flag.BoolVar(&spanish, "s", false, "Use Spanish language.")

	flag.IntVar(&num, "num", 42, "Use 42 language.")
	// flag.Var(&num, "num", "Use 42 language.")

	// why must -name (flag) have an argument unlike -s
	flag.StringVar(&name, "name", "", "A name to say hello to.")

}

func main() {
	flag.Parse()

	fmt.Println(spanish)
	fmt.Println(os.Args[1:])
	fmt.Println(flag.Args())

	if spanish {
		fmt.Printf("Hola %s %d!\n", name, num)

	} else {
		fmt.Printf("Hello %s %d!\n", name, num)
	}

	// default value parameter means the scenario where you don't include the flag at all when invoking the cl app on the cli

	// you can omit the argument of a boolean option(flag) (it would default to true regardless of the default set when specifying the BoolVar() or Bool() method)
	// which makes sense e.g. inputing the flag -s should imply you want spanish (true), you shouldn't have to input s=true. Then
	// if you want false, you input s=false.
	// you dont even have to provide an arg for a boolean option. Indicating the flag without an arg should signify true (as it is in golang).
	// Not indicating/including the option should indicate false (default value declared in source code )

	// -flag x  // non-boolean flags only https://golang.org/pkg/flag/#hdr-Command_line_flag_syntax

	// go run main.go -name Femi -s=true
	// go run main.go -name=Femi -s=false
	// notice below that the string flag -name doesn't have to use the assignment operator to assign its value, and the value is not also
	// recognised as a non-flag CL argument instead it's a flag CL argument. Also, the -s flag
	// go run main.go -name Femi -s=true asdf

	// // go run main.go -name Femi -s true asdf

	// go run main.go -s=true -name=Femi asdf AND go run main.go -s=true -name Femi asdf --- gives the same output at the 3rd println()
	// because a string flag recognizes a following argument as it's argument(flag argument) unlike a non-string? flag which you must
	// use the assignment operator to specify their argument or else the argument is treated as a non-flag argument
	// e.g. go run main.go -name Femi -s true asdf

}
