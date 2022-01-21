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

func NewServer() server {
	sv := server{
		NewJWT(),
		NewDataHandler(),
	}

	return sv
}

func (sv server) Run(port string) {

	r := gin.Default()
	r.POST("/auth", sv.authHandler)
	r.GET("/auth", sv.authGetHandler)
	r.POST("/add", sv.addNoteHandler)
	r.GET("/", sv.getHandler)
	r.GET("/notes/:id", sv.getHandler)
	r.GET("/signout", sv.signoutHandler)
	r.NoRoute(sv.notFoundHandler)

	r.Run(fmt.Sprintf(":" + port))
}

func (sv server) authHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user, err := sv.UserAuth(username, password)

	if err != nil {
		c.JSON(403, gin.H{
			"success": false,
			"error":   "unauthorized",
		})
		return
	}

	token, _ := sv.GenerateJWT(*user)
	c.SetCookie("token", token, 3600, "/", "", true, false)
	c.Redirect(http.StatusFound, "/")
}

func (sv server) authGetHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
}

func (sv server) getHandler(c *gin.Context) {
	user, err := sv.auth(c)
	if err != nil {
		c.JSON(403, gin.H{
			"success": false,
			"error":   "unauthorized",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"user":    user.Username,
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

	title := c.PostForm("title")
	content := c.PostForm("content")

	sv.NoteAdd(user.Id, title, content)
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
		"title":   note.Title,
		"content": note.Content,
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
