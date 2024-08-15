package entities

type MessageId int

type Message struct {
	Id     MessageId
	GameId GameId
	Text   string
	Game   *Game
}

func NewMessage(g *Game, text string) *Message {
	return &Message{
		GameId: g.Id,
		Game:   g,
		Text:   text,
	}
}
