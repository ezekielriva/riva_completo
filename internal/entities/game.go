package entities

type GameId string

type Game struct {
	Id GameId
}

func NewGame() *Game {
	return &Game{}
}
