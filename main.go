package main

import (
	"fmt"
	"log"

	"github.com/Grandeath/Battleship/app"
	"github.com/Grandeath/Battleship/connection"
)

const (
	host = "https://go-pjatk-server.fly.dev"
)

func main() {
	yourTurn := true
	newClient := connection.NewClient(true, host)

	err := app.StartGame(&newClient)
	if err != nil {
		log.Println(err)
	}
	err = app.PrintBoard(&newClient)
	if err != nil {
		log.Println(err)
	}
	err = app.PrintStatus(&newClient)
	if err != nil {
		log.Println(err)
	}

	for yourTurn {
		resp, err := app.FireBullet(&newClient)
		if err != nil {
			log.Println(err)
		}
		if resp == "miss" {
			yourTurn = false
		}
		err = app.PrintBoard(&newClient)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(resp)
		err = app.PrintStatus(&newClient)
		if err != nil {
			log.Println(err)
		}
	}

}
