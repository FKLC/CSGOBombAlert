package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/gen2brain/beeep"
)

type GameState struct {
	Added struct {
		Round struct {
			Bomb bool
		}
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var gameState GameState
		json.NewDecoder(r.Body).Decode(&gameState)
		if gameState.Added.Round.Bomb {
			time.Sleep(30 * time.Second)
			beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
			time.Sleep(4*time.Second + 500*time.Millisecond)
			beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
		}
	})

	ip := flag.String("IP", "127.0.0.1", "IP to listen")
	port := flag.Int("Port", 3000, "Port to listen")
	flag.Parse()

	fmt.Println(fmt.Sprintf("Listening on %s:%d", *ip, *port))
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", *ip, *port), nil)
	if err != nil {
		panic(err)
	}
}
