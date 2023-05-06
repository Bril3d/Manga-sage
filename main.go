package main

import (
	"io"
	"log"
	"manga-sage/controllers"
	"manga-sage/initializers"
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
	// initializers.DB.AutoMigrate(&models.Manga{}, &models.User{}, &models.Chapter{}, &models.Comment{}, &models.Page{}, &models.Rating{})
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
		for {

			Manga := scrapeAresManga()

			for _, m := range Manga {
				for _, u := range m.Urls {
					downloadChapter(u, m.Title)
				}
			}

			// for i := 0; i < len(arr); i++ {
			// 	downloadChapter(arr[i][1], arr[i][0])
			// }

			time.Sleep(time.Minute) // Check once per day
		}

	}()

	// Serve downloaded manga chapters
	r.Static("/mangadownload", "./manga")

	r.Run()
}

type MangaObject struct {
	Title string
	Urls  []string
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
			chapter := t.Find("a")
			new = t.Find("span.status-new").Length() > 0

			url, exists := chapter.Attr("href")

			if exists && new {
				urls = append(urls, url)
			}
		})

		tselection := ts.Find("a .limit .ts-post-image")

		title, exists := tselection.Attr("title")

		if exists && len(urls) > 0 {
			newManga = MangaObject{Title: title, Urls: urls}
			Manga = append(Manga, newManga)
		}

	})

	if err != nil {
		return Manga
	}

	return Manga
}

func downloadChapter(url string, title string) {
	// Download chapter images from AresManga website and store them in local directory
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Get chapter number from URL
	re := regexp.MustCompile(`/*-chapter-(\d+)/`)
	matches := re.FindStringSubmatch(url)
	if len(matches) != 2 {
		return
	}
	chapter := matches[1]
	// Create directory to store chapter images
	err = os.MkdirAll("./manga/"+title+"/"+chapter, 0755)

	if err != nil {
		return
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
			err = downloadImage(src, "./manga/"+title+"/"+chapter+"/image"+strconv.Itoa(i)+".jpg")
			if err != nil {
				log.Println(err)
			}
		}
	})
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


