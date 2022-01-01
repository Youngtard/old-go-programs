package main

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"

	"github.com/abiosoft/ishell"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {
	// play [Song}, shuffle [Directory/Playlist], playlist -shuffle

	shell := ishell.New()

	// ttunes := cli.NewApp()

	// ttunes.Name = "tTunes"
	// ttunes.Usage = "Listen to music from your CLI."

	shell.AddCmd(&ishell.Cmd{
		Name: "play",
		Help: "Play a song",
		Func: func(c *ishell.Context) {
			play("Domaje - Way Down in the Hole.mp3", c)

			// time.Sleep(2 * time.Second)
			// shell.Stop()

		},
	})

	// The action to execute when no subcommands are specified
	// ttunes.Action = func(c *cli.Context) error {
	// 	// fmt.Println("wor")
	// 	shell.Start()
	// 	// time.Sleep(2 * time.Second)
	// 	return nil
	// }

	// ttunes.Run(os.Args)

	shell.Run()

}

func play(songToPlay string, c *ishell.Context) {

	song, _ := os.Open(songToPlay)

	streamer, format, err := mp3.Decode(song)

	if err != nil {
		log.Fatal(err)
	}

	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)

	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		// TODO Remove trailing file extension
		done <- true
	})))

	// c.Printf("Now Playing: %s\n", songToPlay)

	<-done

	// select {}

}
