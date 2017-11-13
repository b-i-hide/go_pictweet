package main

import (
	"database/sql"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	pictweet_db "localhost/pictweet/db"
	"localhost/pictweet/controllers"
)

type Server struct {
	DB *sql.DB
}


func main() {
	db := pictweet_db.InitPictweetDB()
	s := NewServer()
	s.DB = db
	s.Routes()
}

func NewServer() Server {
	return Server{}
}

func (s Server) Routes() {
	router := gin.Default()
	router.HTMLRender = createMyRender()
	loadStaticFiles(router)

	tweets := &controllers.Tweet{DB: s.DB}
	users := &controllers.User{DB: s.DB}

	router.GET("/", tweets.Index)

	router.POST("/users", users.Create)
	router.GET("/users/show/:user_id", users.Show)
	router.GET("/users/new", users.New)

	router.Run(":8080")
}

func createMyRender() multitemplate.Render {
	r := multitemplate.New()

	r.AddFromFiles("tweets_index", "templates/layout/base.html", "templates/tweets/index.html")
	//r.AddFromFiles("tweets_show", "templates/layout/base.html", "templates/tweets/show.html")
	r.AddFromFiles("user_new", "templates/users/new.html")
	r.AddFromFiles("user_show", "templates/layout/base.html", "templates/users/show.html")

	return r
}

func loadStaticFiles(r *gin.Engine) {
	//r.Static("/public/css/", "./public/css")
	//r.Static("/public/js/", "./public/js/")
	//r.Static("/public/fonts/", "./public/fonts/")
	//r.Static("/public/images/", "./public/images/")
}
