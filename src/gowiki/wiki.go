package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

const (
	BASE_DIR = "./"
	TPLE_DIR = BASE_DIR + "tpls/"
	DATA_DIR = BASE_DIR + "data/"
)

type Page struct {
	Title string
	Body []byte
}

var templates = make(map[string]*template.Template)
//var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func init() {
	for _, tpl := range []string{"index", "view", "edit"} {
		templates[tpl] = template.Must(template.ParseFiles(TPLE_DIR + tpl + ".html"))
	}
}

func main() {
	http.HandleFunc("/", makeHandler(indexHandler))
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":7070", nil))
}

func (p *Page) save() error {
	filename := DATA_DIR + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := DATA_DIR + title + ".txt"
	log.Println("filename: " + filename)
	body, err := os.ReadFile(filename)
	//body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tpl string, p *Page)  {
	err := templates[tpl].Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var title string
		rUrlPath := r.URL.Path
		log.Println("rUrlPath:" + rUrlPath)
		if "/" == rUrlPath {
			title = ""
		} else {
			m := validPath.FindStringSubmatch(rUrlPath)
			log.Println(m)
			if m == nil {
				log.Println("url not found!")
				http.NotFound(w, r)
				return
			}
			title = m[2]
		}

		fn(w, r, title)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request, title string)  {
	p := &Page{Title: title}
	renderTemplate(w, "index", p)
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string)  {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/" + title, http.StatusFound)
		return
	}
	//fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	log.Println(r.URL.Path) //`/edit/test`
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{
		Title: title,
		Body: []byte(body),
	}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/" + title, http.StatusFound)
}
