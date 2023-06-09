package main

import (
	"manga-sage/controllers"
	"manga-sage/initializers"
	"manga-sage/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.DB.AutoMigrate(&models.Manga{}, &models.User{}, &models.Chapter{}, &models.Comment{}, &models.Page{}, &models.Rating{})
}

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:5173"} // Replace with your Vue.js frontend URL

	r.Use(cors.New(config))
	r.POST("/manga", controllers.MangaCreate)
	r.GET("/manga", controllers.MangaIndex)
	r.GET("/manga/:id", controllers.MangaShow)
	r.GET("/manga/chapters/:id", controllers.ChaptersShow)
	r.GET("/manga/:id/chapter/:chapter", controllers.PagesShow)
	r.PUT("/manga/:id", controllers.MangaUpdate)
	r.DELETE("/manga/:id", controllers.MangaDelete)
	r.POST("/register", controllers.UserCreate)
	r.POST("/login", controllers.UserLogin)

	//Scrape AresManga website for new manga chapters
	// go func() {
	// 	scrappers.ScrapeMangaLoop()
	// }()

	// Serve downloaded manga chapters
	r.Static("/mangadownload", "./manga")

	r.Run()
}
