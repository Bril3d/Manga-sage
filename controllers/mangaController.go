package controllers

import (
	"manga-sage/initializers"
	"manga-sage/models"

	"github.com/gin-gonic/gin"
)

func MangaCreate(c *gin.Context) {

	var body struct {
		Title       string
		Description string
	}

	c.Bind(&body)
	manga := models.Manga{Title: body.Title, Description: body.Description}
	result := initializers.DB.Create(&manga)
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"manga": manga,
	})
}

func MangaIndex(c *gin.Context) {
	var manga []models.Manga
	initializers.DB.Find(&manga)

	c.JSON(200, gin.H{
		"manga": manga,
	})
}

func MangaShow(c *gin.Context) {
	id := c.Param("id")

	var manga models.Manga

	initializers.DB.Find(&manga, id)

	c.JSON(200, gin.H{
		"manga": manga,
	})
}

func MangaUpdate(c *gin.Context) {
	id := c.Param("id")

	var manga models.Manga

	var body struct {
		Title       string
		Description string
	}
	c.Bind(&body)

	initializers.DB.Find(&manga, id)

	initializers.DB.Model(&manga).Updates(models.Manga{
		Title:       body.Title,
		Description: body.Description,
	})

	c.JSON(200, gin.H{
		"manga": manga,
	})
}

func MangaDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Manga{}, id)

	c.JSON(200, gin.H{
		"message": "deleted successfully",
	})

}
