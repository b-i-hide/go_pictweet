package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Tweet struct {
	DB *sql.DB
}

func (s Tweet) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "tweets_index", gin.H{
		"title": "Index",
	})
}
