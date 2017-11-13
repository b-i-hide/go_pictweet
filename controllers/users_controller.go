package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"localhost/pictweet/models"
	"log"
	"strconv"
)

type User struct {
	DB *sql.DB
}


func (s User) New(c *gin.Context) {
	c.HTML(http.StatusOK, "user_new", gin.H{
		"title": "User registration",
	})
}

func (s User) Create(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password_digest")
	user := models.NewUser()
	user.Name = name
	user.Email = email
	if password == c.PostForm("password_confirmation") {
		user.PasswordDigest = models.HashAndSalt([]byte(password))
		err := user.Insert(s.DB, user)
		if err == nil {
			c.Redirect(301, "/")
		} else {
			// TODO: あとでrenderに変える
			log.Fatalf("error: %s", err)
		}
	} else {
		// TODO: あとでrenderに変える
		c.Redirect(301, "/users/new")
	}
}

func (s User) Show(c *gin.Context) {
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	user := models.NewUser()
	user, err = user.GetByUserId(s.DB, user_id)
	if err == nil {
		c.HTML(http.StatusOK, "user_show", gin.H{
			"title": "User#show",
			"user": user,
		})
	} else {
		log.Fatalf("error: %s", err)
	}
}
