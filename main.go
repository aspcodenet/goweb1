package main

import "github.com/gin-gonic/gin"

type PageView struct {
	CurrentUser string
	PageTitle   string
	Title       string
	Text        string
}

func start(c *gin.Context) {
	c.HTML(200, "home.html", &PageView{Title: "Välkommen", Text: "Tjena moss"})
}

func about(c *gin.Context) {
	//c.HTML(200,"home.html",&PageView{ Title: "Välkommen", Text: "Tjena moss" })
	c.String(200, "About us")
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**")
	router.GET("/", start)
	router.GET("/about", about)
	router.Run(":8080")
}
