package connection

type client interface {
	StartGame() error
	GetBoard() error
}
