package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingController struct {
	db *sql.DB
}

func (h *PingController) Handler(c *gin.Context) {
	var stats sql.DBStats

	if h.db != nil {
		stats = h.db.Stats()
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "pong",
		"db":     stats,
	})
}
