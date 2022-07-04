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
	fmt.Printf("You entered the link \"%s\" and the webpage was successfully processed!", link)
	//fmt.Print("And the webpage content is:")
	//fmt.Printf("%s", pagetext)
	err = ioutil.WriteFile("webpage.html", pagetext, 0644)
	logError(err)
}

func logError(err error) {
	if err != nil {
		log.Fatal("Error encountered!")
	}
}
