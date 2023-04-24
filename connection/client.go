package connection

type ClientInterface interface {
	StartGame() error
	GetBoard() (BoardResp, error)
	GetLongDesc() (DescriptionStruct, error)
	Fire(coordinates string) (fireStructResp, error)
	GetStatus() (StatusStruct, error)
	GetPlayerList() (PlayerListStruct, error)
}
