package main

import (
	"fmt"
	//"io/ioutil"
	"net/http"
	"log"
	//"golang.org/x/net/html"
	//"strings"
	//"bytes"
	//"io"
	//"text/template"
	//"os"
	//"github.com/ConvertAPI/convertapi-go"
	//"github.com/ConvertAPI/convertapi-go/config"
	//"github.com/ConvertAPI/convertapi-go/param"
	//"bytes"
	"strconv"
)

func main() {
	var link string
	//var buf bytes.Buffer
    //w := io.Writer(&buf)
	fmt.Print("Enter the link to process (to process a default link, just press Enter): ")
	fmt.Scanf("%s", &link)
	if link == "" {
		link = "https://www.nairaland.com/7229653/court-orders-upward-review-judges"
	}
	fmt.Println("*********")
	link0 := link
	var page *http.Response
	var pageTrack *http.Response
	var err error
	var pages []*http.Response
	x := 1
	for {
		page, err = http.Get(link)
		logError(err)
		logError(err)
		if pageTrack == nil {
			// do nothing
		} else if page.Request.URL.Path == pageTrack.Request.URL.Path || x == 20 {
			fmt.Println("break!")
			break
		}
		fmt.Println(x, page.Request.URL.Path, link)
		link = link0 + "/" + strconv.Itoa(x)
		pages = append(pages, page)
		pageTrack = page
		x++
	}
	fmt.Println(pages)
}

func logError(err error) {
	if err != nil {
		log.Fatal("Error encountered!", err)
	}
}
