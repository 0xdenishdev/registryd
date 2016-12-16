package main

import (
	"fmt"
	"time"
	events "github.com/gninjava/registryd/eventmanager"
)

func main() {
	for {
		events.Listen()
		fmt.Println("scanning...")
		time.Sleep(5 * time.Second)
	}
}

