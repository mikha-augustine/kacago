package main

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	setupDB()

	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/", mainPage)
	r.GET("/login", loginPage)
	r.POST("/login-ajax", loginAjax)
	r.Run(":8013")
}

func mainPage(c *gin.Context) {
	tmpl, err := template.ParseFiles("main.html")
	if err != nil {
		log.Println("[mainPage] err:", err)
		return
	}
	tmpl.ExecuteTemplate(c.Writer, "main", nil)
}

func loginPage(c *gin.Context) {
	tmpl, err := template.ParseFiles("login.html")
	if err != nil {
		log.Println("[loginPage] err:", err)
		return
	}
	tmpl.ExecuteTemplate(c.Writer, "login", nil)
}

// ============ Login Ajax ============

func loginAjax(c *gin.Context) {
	user := c.PostForm("username")
	pass := c.PostForm("password")

	if isValidUser(user, pass) {

	}
}
