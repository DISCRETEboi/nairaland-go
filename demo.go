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
	err = ioutil.WriteFile("webpage.html", pagetext, 0644)
	logError(err)
	text := string(pagetext)

	doc, err := html.Parse(strings.NewReader(text))
	logError(err)
/*	fmt.Println(doc.Type)
	fmt.Println(doc.Attr)
	fmt.Println(doc.Data)
	fmt.Println(doc.FirstChild)
	fmt.Println(doc.NextSibling)*/
	procNode(doc)

}

func logError(err error) {
	if err != nil {
		log.Fatal("Error encountered!")
	}
}

func procNode(node *html.Node) {
	if node.Type == html.TextNode {
		fmt.Println("-->", node.Data)
	}
	for i := node.FirstChild; i != nil; i = i.NextSibling {
		procNode(i)
	}
}
