## Nairaland - Golang

**This repository stores all the source files involved in creating the Nairaland-Golang application, which re-renders important information from a Nairaland thread in a pdf document, given the thread link. For basic usage, only the following files are needed:**

- **`source.go` - (Golang source code of the application)**  
- **`source.exe` - (Windows executable of the application)**

**The source and executable, with a README, have also been bundled in a zip file `nairaland-go.zip`.**

**The *Golang* executable `source.exe` when launched,  prompts for a Nairaland thread link, and generates a pdf document `new-nairaland-page.pdf` showing the posts on the thread, with some metadata such as the comments, names of commenters, date of creation of thread etc.**

A default link has been provided, which can be used by just pressing *Enter* in the prompt.

Successive outputs are printed to show the stages of the program that were run successfully.

Intermediate html files were also generated. They are only used by the program, but they can still be inspected. They are:

`webpage.html`: existing web page of the Nairaland thread  
`new-nairaland-page.html`: new version of the web page

The *Golang* source code, Windows executable, and the intermediate and final output files all lie in the current working directory.

To run the source code, use the following code outline in the command line:

	CD "directory where the souce code is"
	go run source.go

Note that the *main* program imports the following libraries, which all have to be installed for it to run (*Golang* itself should have been installed too, obviously):

`fmt`, `io/ioutil`, `net/http`, `log`, `golang.org/x/net/html`, `strings`, `bytes`, `io`, `text/template`, `github.com/ConvertAPI/convertapi-go`, `github.com/ConvertAPI/convertapi-go/config`, `github.com/ConvertAPI/convertapi-go/param`, `strconv`

More features will be added in the future.

Take a look at the gallery below ;)

![img1](img1.png)

---

![img2](img2.png)

---

![img3](img3.png)

