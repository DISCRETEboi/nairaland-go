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
	"github.com/ConvertAPI/convertapi-go"
	"github.com/ConvertAPI/convertapi-go/config"
	"github.com/ConvertAPI/convertapi-go/param"
)

var webpage = template.Must(template.New("webpage").Parse(`
<style type="text/css">
  *, *::after, *::before {
    box-sizing: border-box;
  }

  html {
    font-size: 62.5%;
  }

  body {
    font-size: 1.6rem;
    font-family: 'Verdana'!important;
    color: #25283D;
    background-color: #f1f7f9;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
  }

  a {
    color: #428bca;
    text-decoration: none;
  }

 .comment{
   margin-bottom: 20px;
   position: relative;
   z-index: 0;
 }

 .comment .comment-avatar{
   border: 2px solid #fff;
   border-radius: 50%;
   box-shadow: 0 1px 2px rgba(0, 0, 0, .2);
   height: 40px;
   left: 0;
   overflow: hidden;
   position: absolute;
   top: 0;
   width: 40px;
 }

 .comment .comment-avatar img{
   display: block;
   height: 100%;
   width: 100%;
 }

 .comment .comment-box{
   background-color: #fcfcfc;
   border-radius: 4px;
   box-shadow: 0 1px 1px rgba(0, 0, 0, .50);
   margin-left: 50px;
   min-height: 60px;
   position: relative;
   padding: 5px;
   padding-bottom: 5px;
 }

  .comment .comment-main{
   background-color: #fcfcfc;
   border-radius: 4px;
   box-shadow: 0 1px 1px rgba(0, 0, 0, .15);
   margin-left: 0px;
   min-height: 200px;
   position: relative;
   padding: 5px;
 }

  .comment blockquote{
   background-color: #dfe5ed;
   border-radius: 4px;
   box-shadow: 0 1px 1px rgba(0, 0, 0, .50);
   min-height: 60px;
   position: relative;
   padding: 5px;
 }

 .comment .comment-box:before,
 .comment .comment-box:after{
   border-width: 10px 10px 10px 0;
   border-style: solid;
   border-color: transparent #FCFCFC;
   content: "";
   left: -10px;
   position: absolute;
   top: 20px;
 }


 .comment .comment-text{
   color: #555f77;
   font-size: 12px;
   padding-left:10px;
   padding-top: 5px;
 }

 .comment .comment-footer{
   color: #acb4c2;
   font-size: 10px;
   padding-top: 5px;
   padding-left:10px;
 }

  .comment .comment-header{
    line-height:20px;
    font-weight:700;
    font-size:12px;
    padding-bottom:8px;
    padding-top:8px;
    padding-left:10px;
    background-color: #dfe5ed;
 }

 .comment .comment-footer:after{
   content: "";
   display: table;
   clear: both;
 }
</style>

<div class="comment">
  <div class="comment-main">
    <div class="comment-header"><a href="https://htmlcolors.com/user/Antonios" style="color:#428bca">{{.Main.Name}} posted the topic -> {{.Topic}}</a></div>
    <div style="line-height:20px;white-space: pre-wrap;" class="comment-text">{{.Main.Comment}}</div>
    <div class="comment-footer">
      <div class="comment-info">
        <span style="line-height:18px" class="comment-date">{{.Date}} {{.Time}}</span>
      </div>
    </div>
  </div>
</div>

{{range .Posts}}
<div class="comment">
  <div class="comment-avatar">
  	<a href="https://htmlcolors.com/user/Antonios"><img class="mobilecommentimage" style="object-fit:cover" src="https://i.ibb.co/kgHjcs6/change-user.png" border="0"></a>
  </div>

  <div class="comment-box">
    <div class="comment-header"><a href="https://htmlcolors.com/user/Antonios" style="color:#428bca">{{.Name}}</a></div>
    <div style="line-height:20px;white-space: pre-wrap;" class="comment-text">{{.Comment}}</div>
  </div>
</div>
{{end}}

`))

var divs []Post
var ind = 0
var cnt, x, y, z = 1, 1, 1, 1
//var hA = []html.Attribute{html.Attribute{"class","narrow"}}
var elmntcnt = 0
var pageposts Webpage

type Post struct {
	Comment string
	Name string
	//Time string
	//Date string
	//Likes string
}

type Webpage struct {
	Posts []Post
	Main Post
	Topic string
	Time string
	Date string
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
	divs = cleanDivs(divs)
	pageposts.Posts = divs[1:]
	pageposts.Main = divs[0] 
	//fmt.Println(divs)
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
	config.Default.Secret = "Uhy0MidCpF8ZmoUT"
	convertapi.ConvDef("html", "pdf",
		param.NewPath("File", "new-nairaland-page.html", nil)).ToPath("convertapi/new-nairaland-page.pdf")
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
	if node.Type == html.ElementNode  && node.Data == "b" {
		elmntcnt++
	}
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, val := range node.Attr {
			if val.Key == "class" && val.Val == "user" {
				divs = append(divs, Post{"", ""})
				divs[ind].Name = renderNode(node)
				//fmt.Println(divs[ind], ind)
			}
		}
	} else if node.Type == html.ElementNode && node.Data == "div" {
		if node.Attr[0].Key == "class" && node.Attr[0].Val == "narrow" {
			divs = append(divs, Post{"", ""})
			divs[ind].Comment = renderNode(node)
			//fmt.Println(divs[ind], ind)
			ind = ind + 1
		}
	} else if elmntcnt == 8 && x == cnt {
		pageposts.Time = renderNode(node)
		x++
		fmt.Println(pageposts.Time)
	} else if elmntcnt == 9 && y == cnt {
		pageposts.Date = renderNode(node)
		y++
		fmt.Println(pageposts.Date)
	} else if node.Type == html.ElementNode && node.Data == "title" && z == cnt {
		pageposts.Topic = renderNode(node.FirstChild)
		z++
		fmt.Println(pageposts.Topic)
	}
	//cnt++
	for i := node.FirstChild; i != nil; i = i.NextSibling {
		procNode(i)
	}
}

func cleanDivs(divs []Post) []Post {
	pst := Post{}
	for {
		if divs[len(divs)-1] == pst {
			divs = divs[:len(divs)-2]
			//fmt.Println("=====================================================")
		} else {
			break
		}
	}
	return divs
}