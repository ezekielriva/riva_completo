package newgame

import (
	"fmt"

	"github.com/ezekielriva/riva_completo/internal/entities"
	"github.com/ezekielriva/riva_completo/internal/repositories"
)

type NewGameUseCase struct {
	gamesRepo repositories.GamesRepository
}

func NewNewGameUseCase(gamesRepo repositories.GamesRepository) *NewGameUseCase {
	return &NewGameUseCase{
		gamesRepo: gamesRepo,
	}
}

func (u *NewGameUseCase) Execute() (*entities.Game, error) {
	game := u.buildGame()

	err := u.gamesRepo.Create(game)

	if err != nil {
		return nil, fmt.Errorf("execute: %w", err)
	}

	return game, nil
}

func (u *NewGameUseCase) buildGame() *entities.Game {
	return entities.NewGame()
}
