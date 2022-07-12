/*
This package is ultimately designed to take a Nairaland link and output
a pdf document showing the posts, likes, comments etc.
*/
// This current version takes in the link and writes the webpage as an html file

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	"golang.org/x/net/html"
	"strings"
	"bytes"
	"io"
)

var divs []string
//var hA = []html.Attribute{html.Attribute{"class","narrow"}}

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
	fmt.Printf("You entered the link \"%s\" and the webpage was successfully processed!\n", link)
	err = ioutil.WriteFile("webpage.html", pagetext, 0644)
	logError(err)
	text := string(pagetext)
	doc, err := html.Parse(strings.NewReader(text))
	logError(err)
	procNode(doc)
	fmt.Println("+---------------------------------------------------+")
	for i, val := range divs {
		fmt.Println(i)
		fmt.Println(val)
		fmt.Println("+---------------------------------------------------+")
	}
}

func logError(err error) {
	if err != nil {
		log.Fatal("Error encountered!", err)
	}
}

func renderNode(n *html.Node) string {
    var buf bytes.Buffer
    w := io.Writer(&buf)
    html.Render(w, n)
    return buf.String()
}

func procNode(node *html.Node) {
	/*fmt.Println("Level", cnt)
	if node.Type == html.ElementNode {
		fmt.Println("E>", node.Data)
	}
	if node.Type == html.TextNode {
		fmt.Println("T>", node.Data)
	}*/
	if node.Type == html.ElementNode && node.Data == "div" {
		//fmt.Printf("%T\n", renderNode(node))
		if node.Attr[0].Key == "class" && node.Attr[0].Val == "narrow" {
			divs = append(divs, renderNode(node))
		}
	}
	//cnt = cnt + 1
	for i := node.FirstChild; i != nil; i = i.NextSibling {
		procNode(i)
	}
}
