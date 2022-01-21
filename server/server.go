package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type server struct {
	JWTHandler
	DataHandler
}

type formUser struct {
	First_name string `json:"firstname"`
	Last_name  string `json:"lastname"`
	Username   string `json:"user"`
	Password   string `json:"pass"`
}

type formNote struct {
	Id      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewServer(jwt_key, db_host, db_port, db_password string) server {
	sv := server{
		NewJWT(jwt_key),
		NewDataHandler(db_host, db_port, db_password),
	}

	return sv
}

func (sv server) Run(port string) {

	r := gin.Default()
	r.POST("/auth", sv.authHandler)
	r.POST("/reg", sv.addUserHandler)
	r.POST("/add", sv.addNoteHandler)
	r.POST("/update", sv.updateNoteHandler)
	r.GET("/notes/:id", sv.getNoteHandler)
	r.GET("/notes", sv.getUserNotesHandler)
	r.GET("/signout", sv.signoutHandler)
	r.NoRoute(sv.notFoundHandler)

	r.Run(fmt.Sprintf(":" + port))
}

func (sv server) authHandler(c *gin.Context) {
	fu := &formUser{}
	c.BindJSON(&fu)

	user, err := sv.UserAuth(fu.Username, fu.Password)

	if err != nil {
		c.JSON(403, gin.H{
			"success": false,
			"error":   "unauthorized",
		})
		return
	}

	token, _ := sv.GenerateJWT(*user)
	c.SetCookie("token", token, 3600, "/", "", true, false)
	c.JSON(403, gin.H{
		"success": true,
	})
}

func (sv server) addUserHandler(c *gin.Context) {
	fu := &formUser{}
	c.BindJSON(&fu)

	user, err := sv.UserAdd(fu.Username, fu.Password)

	if err != nil {
		c.JSON(403, gin.H{
			"success": false,
			"error":   "unauthorized",
		})
		return
	}

	token, _ := sv.GenerateJWT(*user)
	c.SetCookie("token", token, 3600, "/", "", true, false)

	c.JSON(403, gin.H{
		"success": true,
	})
}

func (sv server) addNoteHandler(c *gin.Context) {
	user, err := sv.auth(c)
	if err != nil {
		c.JSON(403, gin.H{
			"success": false,
			"error":   "unauthorized",
		})
		return
	}

	fn := &formNote{}
	c.BindJSON(&fn)

	sv.NoteAdd(user.Id, fn.Title, fn.Content)
	c.JSON(200, gin.H{
		"success": true,
	})
}

func (sv server) updateNoteHandler(c *gin.Context) {
	user, err := sv.auth(c)
	if err != nil {
		c.JSON(403, gin.H{
			"success": false,
			"error":   "unauthorized",
		})
		return
	}
	fn := &formNote{}
	c.BindJSON(&fn)

	note, err := sv.GetNote(fn.Id)
	if err != nil {
		c.JSON(404, gin.H{
			"success": false,
			"error":   "not found",
		})
		return
	}

	if !user.Admin && user.Id != note.Owner {
		c.JSON(404, gin.H{
			"success": false,
			"error":   "unauthorized",
		})
		return
	}

	sv.NoteUpdate(fn.Id, fn.Title, fn.Content)
	c.JSON(200, gin.H{
		"success": true,
	})
}

func (sv server) getNoteHandler(c *gin.Context) {
	sid := c.Param("id")
	id, err := strconv.ParseUint(sid, 10, 0)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"error":   "bad request",
		})
		return
	}

	user, err := sv.auth(c)
	if err != nil {
		c.JSON(403, gin.H{
			"success": false,
			"error":   "unauthorized",
		})
		return
	}

	note, err := sv.GetNote(uint(id))
	if err != nil {
		c.JSON(404, gin.H{
			"success": false,
			"error":   "not found",
		})
		return
	}

	if !user.Admin && user.Id != note.Owner {
		c.JSON(404, gin.H{
			"success": false,
			"error":   "unauthorized",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"id":      note.Id,
		"title":   note.Title,
		"content": note.Content,
	})
}

func (sv server) getUserNotesHandler(c *gin.Context) {
	user, err := sv.auth(c)
	if err != nil {
		c.JSON(403, gin.H{
			"success": false,
			"error":   "unauthorized",
		})
		return
	}

	notes, err := sv.GetUserNotes(user.Id)
	fmt.Println(notes)
	if err != nil {
		c.JSON(404, gin.H{
			"success": false,
			"error":   "not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"notes":   notes,
	})
}

func (sv server) signoutHandler(c *gin.Context) {
	c.SetCookie("token", "", 0, "/", "", true, false)
	c.Redirect(http.StatusFound, "/")
}

func (sv server) notFoundHandler(c *gin.Context) {
	c.String(404, "404 page not found.")
}

func (sv *server) auth(c *gin.Context) (user User, err error) {
	cookie, err := c.Request.Cookie("token")

	if err != nil {
		return
	}
	user, err = sv.DecodeJWT(cookie.Value)
	return
}
