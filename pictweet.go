package main

import (
	"database/sql"
	db2 "localhost/pictweet/db"
	"github.com/gin-gonic/gin"
)

type Server struct {
	DB *sql.DB
}

func main() {
	db := db2.InitPictweetDB()
	s := NewServer()
	s.DB = db
	s.Routes()
}

func NewServer() Server {
	return Server{}
}

func (s Server) Routes() {
	router := gin.Default()

	router.Run(":3000")
}

//func LoadTemplates(r *gin.Engine) {
//	r.LoadHTMLGlob("templates/**/*")
//}
