package scrappers

import (
	"fmt"
	"log"
	"manga-sage/initializers"
	"manga-sage/models"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type MangaObject struct {
	Title       string
	Urls        []string
	Cover       string
	Description string
	Chapters    []models.Chapter
}

func ScrapeMangaLoop() {

	for {
		Manga := scrapeAresManga("https://aresmanga.net/series/gods-webnovel/")

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
					filepath := "/" + m.Title + "/" + chapter + "/image" + strconv.Itoa(i) + ".webp"
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
						filepath := "./frontend/manga-sage/src/assets/manga/" + m.Title + "/cover.webp"
						downloadImage(m.Cover, filepath)
						Cover_Image := m.Title + "/cover.webp"
						MangaClone := &models.Manga{
							Title:       m.Title,
							Cover_Image: Cover_Image,
							Chapters:    newChapter,
							Description: m.Description,
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
	desc := doc.Find(".entry-content p")

	description := desc.Text()

	tselection := doc.Find(".thumb img")

	title, exists := tselection.Attr("title")

	cover, coverExist := tselection.Attr("src")

	doc.Find("#chapterlist ul li .chbox").Each(func(j int, ts *goquery.Selection) {

		urls = make([]string, 0)

		ts.Find(".eph-num").Each(func(j int, t *goquery.Selection) {
			chapterLink := t.Find("a")
			url, exists := chapterLink.Attr("href")

			if exists {
				urls = append(urls, url)
			}
		})

		if exists && coverExist {
			newManga = MangaObject{Title: title, Urls: urls, Cover: cover, Description: description}
			Manga = append(Manga, newManga)
		}

	})

	if err != nil {
		return Manga
	}

	return Manga
}

func scrapeLatestAresManga(site string) []MangaObject {
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
