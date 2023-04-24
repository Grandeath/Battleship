package app

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Grandeath/Battleship/connection"
)

const (
	host = "https://go-pjatk-server.fly.dev"
)

var NewClient = connection.NewClient(host)

func StartApp() {

	askStartingConditions()

	waitForGame := true

	err := StartGame(&NewClient)
	if err != nil {
		log.Println(err)
	}

	for waitForGame {
		status, _ := GetStatus(&NewClient)
		if status.Game_status == "game_in_progress" {
			waitForGame = false
		}
		time.Sleep(time.Millisecond * 500)
	}
	err = PrintBoard(&NewClient)
	if err != nil {
		log.Println(err)
	}

	err = PrintStatus(&NewClient)
	if err != nil {
		log.Println(err)
	}
	RunApp()
}

func RunApp() {
	gameStillGoint := true
	status := waitFunction()

	for gameStillGoint {
		myTurn()

		status = waitFunction()

		UpdateMyBoard(&NewClient, status.Opp_shots)

		if status.Last_game_status == "ended" || status.Last_game_status == "win" || status.Last_game_status == "lose" {
			fmt.Println(status.Last_game_status)
			gameStillGoint = false
		}
	}
}

func waitFunction() connection.StatusStruct {
	waitingForResponse := true
	gotStatus := connection.StatusStruct{}
	for waitingForResponse {
		status, err := GetStatus(&NewClient)
		if err != nil {
			log.Println(err)
		}
		if status.Should_fire {
			gotStatus = status
			waitingForResponse = false
		}
	}
	time.Sleep(time.Millisecond * 500)
	return gotStatus
}

func myTurn() {
	for {
		resp, err := FireBullet(&NewClient)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if resp == "miss" {
			break
		}
	}
}

func askStartingConditions() {
	scanner := bufio.NewScanner(os.Stdin)
	var question string
	fmt.Println("Do you want to specify nick? (yes/no)")
	if scanner.Scan() {
		question = scanner.Text()
	} else {
		log.Println(scanner.Err())
	}

	if question == "yes" {
		fmt.Println("Write your Nick")

		if scanner.Scan() {
			NewClient.StartingHeader.Nick = scanner.Text()
		} else {
			log.Panicln(scanner.Err())
		}

		fmt.Println("Write your Desc")
		if scanner.Scan() {
			NewClient.StartingHeader.Desc = scanner.Text()
		} else {
			log.Panicln(scanner.Err())
		}
	}

	var secondQuestion string
	fmt.Println("Do you want to play against bot (yes/no)")
	if scanner.Scan() {
		secondQuestion = scanner.Text()
	} else {
		log.Println(scanner.Err())
	}
	if secondQuestion == "yes" {
		NewClient.StartingHeader.Wpbot = true
	} else {
		NewClient.StartingHeader.Wpbot = false
		NewClient.StartingHeader.Target_nick = choosePlayer()
	}
}

func choosePlayer() string {
	playerList, err := GetPlayerList(&NewClient)
	if err != nil {
		log.Println(err)
	}
	if len(playerList.PlayerStruct) == 0 {
		fmt.Println("No enemies to challenge waiting for the oponent")
		return ""
	}
	for index, player := range playerList.PlayerStruct {
		fmt.Printf("%d %s - %s \n", index, player.Nick, player.Game_status)
	}
	fmt.Println("Choose a player number")
	var chosenPlayer string
	_, err = fmt.Scanln(&chosenPlayer)
	if err != nil {
		log.Println(err)
	}
	numPlayer, err := strconv.Atoi(chosenPlayer)
	if err != nil {
		log.Fatal(err)
	}
	if numPlayer >= len(playerList.PlayerStruct) || numPlayer < 0 {
		log.Fatal("num outside the list")
	}

	return playerList.PlayerStruct[numPlayer].Nick
}
