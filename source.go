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

var divs []Post
//var hA = []html.Attribute{html.Attribute{"class","narrow"}}

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
    <div class="comment-header"><a href="https://htmlcolors.com/user/Antonios" style="color:#428bca">A Nairaland User</a></div>
    <div style="line-height:20px;white-space: pre-wrap;" class="comment-text">This is the original post. Check the comments below. [This part will be fixed later]</div>
  </div>
</div>

{{range .Posts}}
<div class="comment">
  <div class="comment-avatar">
  	<a href="https://htmlcolors.com/user/Antonios"><img class="mobilecommentimage" style="object-fit:cover" src="https://i.ibb.co/kgHjcs6/change-user.png" border="0"></a>
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
