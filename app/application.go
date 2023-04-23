package app

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Grandeath/Battleship/connection"
	gui "github.com/grupawp/warships-lightgui"
)

var NewBoard = gui.New(
	gui.NewConfig())

func StartGame(client connection.ClientInterface) error {
	err := client.StartGame()
	if err != nil {
		return err
	}

	return nil
}

func PrintBoard(client connection.ClientInterface) error {

	gotBoard, err := client.GetBoard()
	if err != nil {
		return err
	}

	NewBoard.Import(gotBoard.Board)
	NewBoard.Display()

	return nil
}

func PrintStatus(client connection.ClientInterface) error {
	statusMessage, err := client.GetStatus()
	if err != nil {
		return err
	}
	fmt.Println(statusMessage.Nick)
	fmt.Println(statusMessage.Desc)

	fmt.Println(statusMessage.Opponent)
	fmt.Println(statusMessage.Opp_desc)

	return nil
}

func FireBullet(client connection.ClientInterface) (string, error) {
	fmt.Println("Write coordinated to shot")

	var coord string

	_, err := fmt.Scanln(&coord)
	if err != nil {
		return "", err
	}
	if len(coord) > 3 {
		return "", fmt.Errorf("wrong ammount of characters")
	}
	coord = strings.ToUpper(coord)
	_, ok := BoardMap[coord[0]]
	if !ok {
		return "", fmt.Errorf("wrong first coordinate")
	}

	secondCoord, err := strconv.Atoi(string(coord[1:]))
	if err != nil {
		return "", fmt.Errorf("second coordinate is not a number")
	}
	if secondCoord < 1 || secondCoord > 10 {
		return "", fmt.Errorf("wrong second coordinate")
	}

	resp, err := client.Fire(coord)
	if err != nil {
		return "", err
	}

	fmt.Println(resp.Result)
	return resp.Result, nil
}
