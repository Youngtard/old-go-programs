package main

import (
	"strings"
	"time"

	"github.com/abiosoft/ishell"

	_ "github.com/manifoldco/promptui"
)

func main() {

	// enter ttunes shell
	// play or playlist sub shell/prompt

	// create new shell.
	// by default, new shell includes 'exit', 'help' and 'clear' commands.
	shell := ishell.New()

	// display welcome info.
	shell.Println("Sample Interactive Shell")

	// register a function for "greet" command.
	shell.AddCmd(&ishell.Cmd{
		Name: "greet",
		Help: "greet user",
		Func: func(c *ishell.Context) {
			c.Println("Hello", strings.Join(c.Args, " "))
		},
	})

	// run shell
	// shell.Run()
	shell.Start()
	// shell.Wait()

	time.Sleep(2 * time.Second)
	shell.Stop()
	// validate := func(input string) error {
	// 	_, err := strconv.ParseFloat(input, 64)

	// 	if err != nil {
	// 		return errors.New("Invalid number")
	// 	}

	// 	return nil
	// }

	// prompt := promptui.Prompt{
	// 	Label:    "Number",
	// 	Validate: validate,
	// }

	// result, err := prompt.Run()

	// if err != nil {
	// 	fmt.Printf("Prompt failed %v\n", err)
	// 	return
	// }

	// fmt.Printf("You choose %q\n", result)

	// prompt := promptui.Select{
	// 	Label: "Select Day",
	// 	Items: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday",
	// 		"Saturday", "Sunday"},
	// }

	// _, result, err := prompt.Run()

	// if err != nil {
	// 	fmt.Printf("Prompt failed %v\n", err)
	// 	return
	// }

	// fmt.Printf("You choose %q\n", result)
}
