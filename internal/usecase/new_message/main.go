package newmessage

import (
	"github.com/ezekielriva/riva_completo/internal/entities"
	"github.com/ezekielriva/riva_completo/internal/repositories"
)

type NewMessageUseCase struct {
	Game              *entities.Game
	MessageRepository repositories.MessagesRepository
}

func NewNewMessageUseCase(game *entities.Game, msgRepo repositories.MessagesRepository) *NewMessageUseCase {
	return &NewMessageUseCase{
		Game:              game,
		MessageRepository: msgRepo,
	}
}

// Receives a user input and return a AI generated message based on it
func (u *NewMessageUseCase) Execute(text string) (*entities.Message, error) {
	usrMsg := u.buildMessage(text)
	u.createMessage(usrMsg)

	resMsg := u.buildMessage("moving right")
	u.createMessage(resMsg)

	return resMsg, nil
}

func (u *NewMessageUseCase) buildMessage(text string) *entities.Message {
	return entities.NewMessage(u.Game, text)
}

func (u *NewMessageUseCase) createMessage(m *entities.Message) error {
	return u.MessageRepository.Create(m)
}
