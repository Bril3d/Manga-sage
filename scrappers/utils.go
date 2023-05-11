package scrappers

import (
	"fmt"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/chai2010/webp"
)

func downloadChapter(url string, title string) (string, int) {
	ChapterLength := 0
	// Download chapter images from websites and store them in local directory
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

	// Find all img elements inside the chapter
	doc.Find("#readerarea p img").Each(func(i int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		if exists {
			// Download the image and save it to a file
			filepath := "./frontend/manga-sage/src/assets/manga/" + title + "/" + chapter + "/image" + strconv.Itoa(i) + ".webp"
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

	// Download the image
	resp, err := http.Get(parts[0])
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Decode the image
	img, err := jpeg.Decode(resp.Body)
	if err != nil {
		return err
	}

	// Encode the image in WebP format
	options := webp.Options{Lossless: false, Quality: 80}
	err = webp.Encode(file, img, &options)
	if err != nil {
		return err
	}

	return nil
}
