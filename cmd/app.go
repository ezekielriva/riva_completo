package main

import (
	"github.com/ezekielriva/riva_completo/internal/controllers"
	"github.com/ezekielriva/riva_completo/internal/database"
)

func main() {
	db := database.InitDB()
	
	r := controllers.SetupRouter(db)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
