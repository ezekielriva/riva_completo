package repositories

import "github.com/ezekielriva/riva_completo/internal/entities"

type GamesRepository interface {
	// Creates a record in the database
	// Returns the game argument modified with database data.
	Create(u *entities.Game) error
}
