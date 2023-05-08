package scrappers

import (
	"fmt"
	"io"
	"log"
	"manga-sage/initializers"
	"manga-sage/models"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type MangaObject struct {
	Title    string
	Urls     []string
	Cover    string
	Chapters []models.Chapter
}

func ScrapeMangaLoop() {

	for {
		Manga := scrapeAresManga("https://aresmanga.net/")
		fmt.Println(Manga)
		for _, m := range Manga {
			var newChapter []models.Chapter

			var chapter string
			var ChapterLength int
			var existingManga models.Manga
			for _, u := range m.Urls {
				var pages []models.Page

				chapter, ChapterLength = downloadChapter(u, m.Title)
				// Create a new Page struct for each page and append it to the slice
				for i := 0; i < ChapterLength; i++ {
					filepath := "/" + m.Title + "/" + chapter + "/image" + strconv.Itoa(i) + ".jpg"
					page := models.Page{
						Image: filepath,
					}

					pages = append(pages, page)
				}
				// Check if the chapter already exists in the database
				var existingChapter models.Chapter
				result := initializers.DB.Where("title = ? AND number = ?", m.Title, chapter).First(&existingChapter)

				if result.Error != nil {
					DBresult := initializers.DB.Where("title = ?", m.Title).First(&existingManga)
					if DBresult.Error != nil {
						filepath := "./frontend/manga-sage/src/assets/manga/" + m.Title + "/cover.jpg"
						downloadImage(m.Cover, filepath)
						Cover_Image := m.Title + "/cover.jpg"
						MangaClone := &models.Manga{
							Title:       m.Title,
							Cover_Image: Cover_Image,
							Chapters:    newChapter,
						}

						result := initializers.DB.Create(MangaClone)

						if result.Error != nil {
							log.Fatal(result.Error)
						}
					} else {

						// Manga already exists in database, do nothing
						log.Printf("Manga '%s' already exists in database, skipping", m.Title)
					}
					initializers.DB.Where("title = ?", m.Title).First(&existingManga)

					// Chapter doesn't exist in the database, create a new one
					newChapter := &models.Chapter{
						Pages:   pages,
						Number:  chapter,
						MangaID: existingManga.ID,
					}

					initializers.DB.Create(newChapter)

				} else {
					// Chapter already exists in the database, skip creating a new one
					log.Printf("Chapter '%s' of manga '%s' already exists in database, skipping", chapter, m.Title)
				}

			}

		}

		time.Sleep(30 * time.Minute) // Check once per half an hour
	}
}

func scrapeAresManga(site string) []MangaObject {
	var Manga []MangaObject
	var urls []string
	// Use goquery to scrape AresManga website and extract URLs for new chapters
	doc, err := goquery.NewDocument(site)

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
	err = createDirectoryIfNotExist("./frontend/manga-sage/src/assets/manga/" + title + "/" + chapter)

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
			filepath := "./frontend/manga-sage/src/assets/manga/" + title + "/" + chapter + "/image" + strconv.Itoa(i) + ".jpg"
			ChapterLength++
			err = downloadImage(src, filepath)

			if err != nil {
				log.Println(err)
			}
		}
	})

	return chapter, ChapterLength
}

func createDirectoryIfNotExist(dir string) error {
	_, err := os.Stat(string(dir))

	if err == nil {
		fmt.Println("directory exists " + dir)
		// directory exists, do nothing
		return nil
	}
	if os.IsNotExist(err) {
		fmt.Println("directory doesnt exists " + dir)
		// directory does not exist, create it
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	} else {
		// unexpected error occurred while checking if directory exists
		return err
	}
	return nil
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
	parts := strings.Split(url, "?")
	// Download the image and write it to the file
	resp, err := http.Get(parts[0])

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
