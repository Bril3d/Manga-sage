package main

import (
	"fmt"
	"io"
	"log"
	"manga-sage/controllers"
	"manga-sage/initializers"
	"manga-sage/models"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
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
	config.AllowOrigins = []string{"http://127.0.0.1:5173"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowCredentials = true
	r.Use(cors.New(config))
	r.POST("/manga", controllers.MangaCreate)
	r.GET("/manga", controllers.MangaIndex)
	r.GET("/manga/:id", controllers.MangaShow)
	r.PUT("/manga/:id", controllers.MangaUpdate)
	r.DELETE("/manga/:id", controllers.MangaDelete)
	r.POST("/register", controllers.UserCreate)
	r.POST("/login", controllers.UserLogin)

	// Scrape AresManga website for new manga chapters
	go func() {
		var newChapter []models.Chapter

		for {
			Manga := scrapeAresManga()

			for _, m := range Manga {

				var chapter string
				var ChapterLength int

				for _, u := range m.Urls {
					var pages []models.Page

					chapter, ChapterLength = downloadChapter(u, m.Title)
					// Create a new Page struct for each page and append it to the slice
					for i := 0; i < ChapterLength; i++ {
						filepath := "./manga/" + m.Title + "/" + chapter + "/image" + strconv.Itoa(i) + ".jpg"
						page := models.Page{
							Image: filepath,
						}

						pages = append(pages, page)
					}

					newChapter = append(newChapter, models.Chapter{
						Pages:  pages,
						Number: chapter,
					})
				}

				filepath := "./manga/" + m.Title + "/cover.jpg"
				downloadImage(m.Cover, filepath)

				MangaClone := &models.Manga{
					Title:       m.Title,
					Cover_Image: filepath,
					Chapters:    newChapter,
				}

				result := initializers.DB.Create(MangaClone)

				if result.Error != nil {
					log.Fatal(result.Error)
				}
			}

			time.Sleep(30 * time.Minute) // Check once per half an hour
		}
	}()

	// Serve downloaded manga chapters
	r.Static("/mangadownload", "./manga")

	r.Run()
}

type MangaObject struct {
	Title    string
	Urls     []string
	Cover    string
	Chapters []models.Chapter
}

func scrapeAresManga() []MangaObject {
	var Manga []MangaObject
	var urls []string
	// Use goquery to scrape AresManga website and extract URLs for new chapters
	doc, err := goquery.NewDocument("https://aresmanga.net/")

	var newManga MangaObject
	var new bool

	doc.Find(".styletere .bsx").Each(func(j int, ts *goquery.Selection) {

		urls = make([]string, 0)

		ts.Find(".bigor .chfiv li").Each(func(j int, t *goquery.Selection) {
			chapterLink := t.Find("a")
			new = t.Find("span.status-new").Length() > 0

			url, exists := chapterLink.Attr("href")

			if exists && new {
				urls = append(urls, url)
			}
		})

		tselection := ts.Find("a .limit .ts-post-image")

		title, exists := tselection.Attr("title")

		cover, coverExist := tselection.Attr("src")

		if exists && coverExist && new {
			newManga = MangaObject{Title: title, Urls: urls, Cover: cover}
			Manga = append(Manga, newManga)
		}

	})

	if err != nil {
		return Manga
	}

	return Manga
}

func downloadChapter(url string, title string) (string, int) {
	ChapterLength := 0
	// Download chapter images from AresManga website and store them in local directory
	resp, err := http.Get(url)
	if err != nil {
		return "", 0
	}
	defer resp.Body.Close()

	chapter := getChapterNumber(url)

	// Create directory to store chapter images
	err = os.MkdirAll("./manga/"+title+"/"+chapter, 0755)

	if err != nil {
		return "", 0
	}

	// Parse HTML response
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	// Find all img elements inside #readerarea and extract the src attribute value
	doc.Find("#readerarea p img").Each(func(i int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		if exists {
			// Download the image and save it to a file
			filepath := "./manga/" + title + "/" + chapter + "/image" + strconv.Itoa(i) + ".jpg"
			ChapterLength++
			err = downloadImage(src, filepath)

			if err != nil {
				log.Println(err)
			}
		}
	})

	return chapter, ChapterLength
}
func getChapterNumber(url string) string {
	// Get chapter number from URL
	re := regexp.MustCompile(`/*-chapter-(\d+)/`)
	matches := re.FindStringSubmatch(url)
	if len(matches) != 2 {
		return ""
	}
	return matches[1]
}
func downloadImage(url string, filename string) error {
	// Create the file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Download the image and write it to the file
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
