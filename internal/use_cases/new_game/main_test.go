package newgame

import (
	"fmt"
	"testing"

	"github.com/ezekielriva/riva_completo/internal/entities"
	mock_repositories "github.com/ezekielriva/riva_completo/internal/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewGameUseCaseHappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	gamesRepo := mock_repositories.NewMockGamesRepository(ctrl)
	gamesRepo.EXPECT().Create(gomock.Any()).DoAndReturn(func(u *entities.Game) error {
		u.Id = "1234"
		return nil
	})

	game, err := NewNewGameUseCase(gamesRepo).Execute()

	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, game, "Game must be present")
	assert.NotEmpty(t, game.Id, "GameID must be set")
	assert.Equal(t, string(game.Id), "1234", "GameID must match")
}

func TestNewGameUseCaseRepositoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	gamesRepo := mock_repositories.NewMockGamesRepository(ctrl)
	gamesRepo.EXPECT().Create(gomock.Any()).DoAndReturn(func(u *entities.Game) error {
		return fmt.Errorf("unable to save")
	})

	game, err := NewNewGameUseCase(gamesRepo).Execute()

	assert.Empty(t, game, "Game shoudln't be returned")
	assert.ErrorContains(t, err, "execute: unable to save")
}
