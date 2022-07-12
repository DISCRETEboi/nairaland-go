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
	if link == "" {
		link = "https://www.nairaland.com/7217386/strike-fg-breaks-asuus-rank"
	}
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
	var cnt = 1
	procNode(doc, cnt)

}

func logError(err error) {
	if err != nil {
		log.Fatal("Error encountered!")
	}
}

func procNode(node *html.Node, cnt int) {
	fmt.Println("Level", cnt)
	if node.Type == html.ElementNode {
		fmt.Println("E>", node.Data)
	}
	if node.Type == html.TextNode {
		fmt.Println("T>", node.Data)
	}
	/*if node.Type == html.ElementNode && node.Data == "div" {
		fmt.Println(node.FirstChild.Data)
	}*/
	//cnt = cnt + 1
	for i := node.FirstChild; i != nil; i = i.NextSibling {
		procNode(i, cnt+1)
	}
}



