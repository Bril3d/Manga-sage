package controllers

import (
	"log"
	"manga-sage/initializers"
	"manga-sage/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	page := c.DefaultQuery("page", "1") // Get the page number from the query parameter
	limit := 16                         // Number of items per page

	var count int64
	if err := initializers.DB.Model(&models.Manga{}).Count(&count).Error; err != nil {
		log.Fatal(err)
	}

	var offset int64
	pageNum, err := strconv.ParseInt(page, 10, 64)
	if err != nil || pageNum < 1 {
		offset = 0
	} else {
		offset = (pageNum - 1) * int64(limit) // Convert limit to int64
	}

	var mangas []models.Manga

	err = initializers.DB.Offset(int(offset)).Limit(int(limit)).Order("id DESC").Find(&mangas).Error // Convert offset and limit to int
	if err != nil {
		log.Fatal(err)
	}

	initializers.DB.Model(&models.Manga{}).Preload("Chapters", func(tx *gorm.DB) *gorm.DB {
		return tx.Joins(`JOIN LATERAL (
				SELECT c.id FROM Chapters c WHERE c.manga_id = chapters.manga_id ORDER BY c.manga_id DESC LIMIT 2
				) AS ch ON ch.id = chapters.id`)
	}).Find(&mangas)

	totalPages := int(math.Ceil(float64(count) / float64(limit)))

	c.JSON(200, gin.H{
		"page":       pageNum,
		"per_page":   limit,
		"total":      count,
		"totalPages": totalPages,
		"manga":      mangas,
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
