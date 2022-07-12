package main

import (
    "fmt"
    "os"
    "golang.org/x/net/html"
    "net/http"
    "log"
    "strings"
    "io"
)

func main() {
    //if the caller didn't provide a URL to fetch...
    if len(os.Args) < 2 {
        //print the usage and exit with an error
        fmt.Printf("usage:\n  pagetitle <url>\n")
        os.Exit(1)
    }

    URL := os.Args[1]

    //GET the URL
	resp, err := http.Get(URL)

	//if there was an error, report it and exit
	if err != nil {
    	//.Fatalf() prints the error and exits the process
    	log.Fatalf("error fetching URL: %v\n", err)
	}

	//make sure the response body gets closed
	defer resp.Body.Close()

	//check response status code
	if resp.StatusCode != http.StatusOK {
    	log.Fatalf("response status code was %d\n", resp.StatusCode)
	}

	//check response content type
	ctype := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(ctype, "text/html") {
    	log.Fatalf("response content type was %s not text/html\n", ctype)
	}

	//create a new tokenizer over the response body
	tokenizer := html.NewTokenizer(resp.Body)

	//loop until we find the title element and its content
	//or encounter an error (which includes the end of the stream)
	for {
    	//get the next token type
    	tokenType := tokenizer.Next()

    	//if it's an error token, we either reached
    	//the end of the file, or the HTML was malformed
    	if tokenType == html.ErrorToken {
        	err := tokenizer.Err()
        	if err == io.EOF {
            	//end of the file, break out of the loop
            	break
        	}
        	//otherwise, there was an error tokenizing,
        	//which likely means the HTML was malformed.
        	//since this is a simple command-line utility,
        	//we can just use log.Fatalf() to report the error
        	//and exit the process with a non-zero status code
        	log.Fatalf("error tokenizing HTML: %v", tokenizer.Err())
    	}

    	//process the token according to the token type...
    	//if this is a start tag token...
		if tokenType == html.StartTagToken {
    		//get the token
    		token := tokenizer.Token()
    		//if the name of the element is "title"
    		if "div" == token.Data {
        		//the next token should be the page title
        		tokenType = tokenizer.Next()
        		//just make sure it's actually a text token
        		if tokenType == html.TextToken {
            		//report the page title and break out of the loop
            		fmt.Println(tokenizer.Token().Data)
            		break
        		}
    		}
		}
	}
}

