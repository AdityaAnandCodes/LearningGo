
package main

import (
	"net/http"
	"os"
	"strings"
	"log"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	url:= "https://en.wikipedia.org/wiki/Bhutan"

	resp,err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching URL:", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err!= nil {
		log.Fatal("Error parsing HTML:", err)
	}

	file, err := os.Create("output.txt")
	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer file.Close()

	doc.Find("h1,h2,h3,h4,h5,h6").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if text != "" {
			file.WriteString("Heading: " + text +"\n")
	}})

}
