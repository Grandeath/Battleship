package connection

type ClientInterface interface {
	StartGame() error
	GetBoard() (BoardResp, error)
	GetStatus() (statusStruct, error)
	Fire(coordinates string) (fireStructResp, error)
}
