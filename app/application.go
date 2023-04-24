package app

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Grandeath/Battleship/connection"
	gui "github.com/grupawp/warships-lightgui/v2"
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
	statusMessage, err := client.GetLongDesc()
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
	fmt.Println("Write coordinates to shot")

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

	state := gui.Miss
	if resp.Result == "hit" {
		state = gui.Hit
	}

	NewBoard.Set(gui.Right, coord, state)
	NewBoard.Display()

	fmt.Println(resp.Result)
	return resp.Result, nil
}

func GetStatus(client connection.ClientInterface) (connection.StatusStruct, error) {
	return client.GetStatus()
}

func UpdateMyBoard(client connection.ClientInterface, coords []string) error {
	for _, coord := range coords {
		status, err := NewBoard.HitOrMiss(gui.Left, coord)
		if err != nil {
			return err
		}
		NewBoard.Display()
		fmt.Println(status)
		fmt.Println(coord)

	}
	return nil
}

func GetPlayerList(client connection.ClientInterface) (connection.PlayerListStruct, error) {
	return client.GetPlayerList()
}
