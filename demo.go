package main

import (
	"fmt"
	"io/ioutil"
	"golang.org/x/net/html"
	"net/http"
	"log"
	"strings"
)

func main() {
	var link string
	fmt.Print("Enter the link to process: ")
	fmt.Scanf("%s", &link)
	page, err := http.Get(link)
	logError(err)
	pagetext, err := ioutil.ReadAll(page.Body)
	logError(err)
	page.Body.Close()
	text := string(pagetext)

	doc, err := html.Parse(strings.NewReader(text))
	logError(err)
	fmt.Println(doc.Type)
/*	fmt.Println(html.DocumentNode)
	fmt.Println(html.TextNode)
	fmt.Println(html.ElementNode)
	fmt.Println(html.ErrorNode)
*/
}

func logError(err error) {
	if err != nil {
		log.Fatal("Error encountered!")
	}
}
