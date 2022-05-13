package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
//    "path/filepath"
)

type Page struct {
    Title string
	Raw string
	Search string
    Body template.HTML
}

var linksRe = regexp.MustCompile("\\[[a-zA-Z]+\\]")

func loadPage(title string) (*Page, error) {
	raw, err := os.ReadFile("data/" + title + ".txt")
	if err != nil {
		return nil, err
	}
	body := template.HTMLEscapeString(string(raw))
	body = linksRe.ReplaceAllStringFunc(body, func(s string) string {
		m := s[1 : len(s)-1]
		return fmt.Sprintf("<a href=\"/view/%s\">%s</a>", m, m)
	})
	return &Page{Title: title, Raw: string(raw), Body: template.HTML(body)}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
    p, err := loadPage(title)
	if err != nil {
        //http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }
	renderTemplate(w, "home", p)
}

func viewHomeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "home.html", http.StatusFound)
}

var templates = template.Must(template.ParseFiles("tmpl/home.html"))
var validPath = regexp.MustCompile("^/(home)/([a-zA-Z0-9]+)$")

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	//http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/home", viewHomeHandler)

	log.Println("starting server")
	log.Fatal(http.ListenAndServe(":8081", nil))
}