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
	"text/template"
	//"os"
	//"convertapi-go"
)

var divs []Post
//var hA = []html.Attribute{html.Attribute{"class","narrow"}}

var webpage = template.Must(template.New("webpage").Parse(`
<link rel="stylesheet" href="./Google Color Picker _ Html Colors_files/style.css"> <!-- Resource style -->
<link rel="stylesheet" href="./Google Color Picker _ Html Colors_files/font-awesome.min.css">

<div class="comment">
  <div class="comment-main">
    <div class="comment-header"><a href="https://htmlcolors.com/user/Antonios" style="color:#428bca">A Nairaland User</a></div>
    <div style="line-height:20px;white-space: pre-wrap;" class="comment-text">This is the original post. Check the comments below. [This part will be fixed later]</div>
  </div>
</div>

{{range .Posts}}
<div class="comment">
  <div class="comment-avatar">
		<a href="https://htmlcolors.com/user/Antonios"><img class="mobilecommentimage" style="object-fit:cover" src="./Google Color Picker _ Html Colors_files/change-user.png"></a>
  </div>

  <div class="comment-box">
    <div class="comment-header"><a href="https://htmlcolors.com/user/Antonios" style="color:#428bca">A Nairaland User</a></div>
    <div style="line-height:20px;white-space: pre-wrap;" class="comment-text">{{.Comment}}</div>
    <div class="comment-footer">
      <div class="comment-info">
        <span style="line-height:18px" class="comment-date">2019-07-04 09:22:48</span>
      </div>

    </div>
  </div>
</div>
{{end}}

`))

type Post struct {
	Comment string
	//Name string
	//Time string
}

type Webpage struct {
	Posts []Post
}

func main() {
	var link string
	var buf bytes.Buffer
    w := io.Writer(&buf)
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
	pageposts := Webpage{divs}
	//fmt.Println(pageposts)
	err = webpage.Execute(w, pageposts)
	logError(err)
	err = ioutil.WriteFile("new-nairaland-page.html", []byte(buf.String()), 0644)
	logError(err)
	/*fmt.Println("+---------------------------------------------------+")
	for i, val := range divs {
		fmt.Println(i)
		fmt.Println(val)
		fmt.Println("+---------------------------------------------------+")
	}*/
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
			divs = append(divs, Post{renderNode(node)})
		}
	}
	//cnt = cnt + 1
	for i := node.FirstChild; i != nil; i = i.NextSibling {
		procNode(i)
	}
}
