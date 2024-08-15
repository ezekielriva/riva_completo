package repositories

import "github.com/ezekielriva/riva_completo/internal/entities"

type MessagesRepository interface {
	// Creates a record in the database
	// Returns the game argument modified with database data.
	Create(m *entities.Message) error
}
