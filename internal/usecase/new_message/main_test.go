package newmessage

import (
	"testing"

	"github.com/ezekielriva/riva_completo/internal/entities"
	mock_repositories "github.com/ezekielriva/riva_completo/internal/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewMessageUseCaseHappyPath(t *testing.T) {
	game := entities.NewGame()
	game.Id = "1234"
	userInputText := "move right"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	msgRepo := mock_repositories.NewMockMessagesRepository(ctrl)

	msgRepo.EXPECT().Create(&entities.Message{
		GameId: game.Id,
		Text:   userInputText,
		Game:   game,
	}).DoAndReturn(func(m *entities.Message) error {
		m.Id = 1
		return nil
	})

	msgRepo.EXPECT().Create(&entities.Message{
		GameId: game.Id,
		Text:   "moving right",
		Game:   game,
	}).DoAndReturn(func(m *entities.Message) error {
		m.Id = 2
		return nil
	})

	resMsg, err := NewNewMessageUseCase(game, msgRepo).Execute(userInputText)

	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, resMsg, "Message must be present")
	assert.NotEmpty(t, resMsg.Id, "MessageId must be set")
	assert.Equal(t, resMsg.GameId, game.Id, "Message must be related to a game")
	assert.Equal(t, resMsg.Text, "moving right", "Message must contains a text")
}
