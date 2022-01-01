package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var c *cli.Context

func main() {

	// option or global argument/parameter
	var (
		helloCmd cli.ActionFunc
	)

	helloCmd = func(c *cli.Context) error {
		if c.String("name") != "" {
			// c.Args
			fmt.Printf("Hello %s\n", c.String("name"))

		} else {
			fmt.Printf("Hello world")

		}
		return nil
	}

	ttunes := cli.NewApp()

	ttunes.Name = "tTunes"
	ttunes.Usage = "Listen to music in your terminal."

	// The action to execute when no subcommands are specified
	ttunes.Action = func(c *cli.Context) error {
		fmt.Printf("Play! I say %s", c.Args().Get(0))
		return nil
	}

	// Global flags
	ttunes.Flags = []cli.Flag{
		cli.BoolTFlag{
			Name: "lang, l",
		},
	}

	// Hello cmd flags
	helloFlags := []cli.Flag{
		cli.StringFlag{
			Name: "name, n",
		},
	}

	ttunes.Commands = []cli.Command{
		{
			Name:   "hello",
			Usage:  "Say Hello",
			Action: helloCmd,
			Flags:  helloFlags,
		},

		{},
	}

	for {
		err := ttunes.Run(os.Args)

		if err != nil {
			log.Fatal(err)
		}

	}

	// song, _ := os.Open("Airboy_-_Ayepo__prod_Phantom_.mp3")

	// streamer, format, err := mp3.Decode(song)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer streamer.Close()

	// speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// speaker.Play(streamer)

	// select {}

}
