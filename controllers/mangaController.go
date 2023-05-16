package controllers

import (
	"log"
	"manga-sage/initializers"
	"manga-sage/models"
	"net/http"

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
	var mangas []models.Manga
	err := initializers.DB.Order("id DESC").Find(&mangas).Error
	if err != nil {
		log.Fatal(err)
	}

	for i := range mangas {
		var chapters []models.Chapter
		err := initializers.DB.
			Where("manga_id = ?", mangas[i].ID).
			Order("number DESC").
			Limit(2).
			Find(&chapters).Error
		if err != nil {
			log.Fatal(err)
		}
		mangas[i].Chapters = chapters
	}

	c.JSON(200, gin.H{
		"manga": mangas,
	})
}

func ChaptersShow(c *gin.Context) {
	id := c.Param("id")

	type Chapter struct {
		Number string
	}

	var chapters []Chapter

	err := initializers.DB.Select("Number").Where("manga_id = ?", id).Find(&chapters).Error

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"chapters": chapters,
	})
}

func PagesShow(c *gin.Context) {
	mangaID := c.Param("id")
	chapterID := c.Param("chapter")

	var chapter models.Chapter
	if err := initializers.DB.Preload("Pages").Where("Number = ? AND manga_id = ?", chapterID, mangaID).First(&chapter).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Chapter not found"})
		return
	}

	c.JSON(http.StatusOK, chapter.Pages)
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
