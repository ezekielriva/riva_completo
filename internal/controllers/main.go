package controllers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {

	r := gin.Default()

	controller := PingController{db: db}

	r.GET("/ping", controller.Handler)

	return r
}
