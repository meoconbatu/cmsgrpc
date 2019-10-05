package view

import (
	"html/template"
	"net/http"
	"strings"
)

// Tmpl is a reference to all of our templates
var Tmpl = template.Must(template.ParseGlob("templates/*"))

// ServeIndex serve Index
func ServeIndex(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "Go Project CMS",
		Content: "Welcome to our homepage!",
		// Posts: []*Post{
		// 	&Post{
		// 		Title:         "Hello, World!",
		// 		Content:       "Hello world! Thank for comming to the site.",
		// 		DatePublished: time.Now(),
		// 	},
		// 	&Post{
		// 		Title:         "A Post with Comments",
		// 		Content:       "Here is a controversial post. It's sure to attract comments.",
		// 		DatePublished: time.Now().Add(-time.Hour),
		// 		Comments: []*Comment{
		// 			&Comment{
		// 				Author:        "Ben Tranter",
		// 				Comment:       "Nevermind, I guess I just commented on my own post.",
		// 				DatePublished: time.Now().Add(-time.Hour / 2),
		// 			},
		// 		},
		// 	},
		// },
	}
	Tmpl.ExecuteTemplate(w, "page", p)
}

// ServePage serves a page basedd on the route matched. This will match  any URL beginning with /page
func ServePage(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimLeft(r.URL.Path, "/page/")
	if path == "" {
		pages, err := GetPages()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		Tmpl.ExecuteTemplate(w, "pages", pages)
		return
	}
	page, err := GetPage(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	Tmpl.ExecuteTemplate(w, "page", page)
}

// ServePost serves a post
// func ServePost(w http.ResponseWriter, r *http.Request) {
// 	path := strings.TrimLeft(r.URL.Path, "/post/")
// 	if path == "" {
// 		http.NotFound(w, r)
// 		return
// 	}
// 	p := &Post{
// 		Title:   strings.ToTitle(path),
// 		Content: "Here is my post",
// 		Comments: []*Comment{
// 			&Comment{
// 				Author:        "Ben Tranter",
// 				Comment:       "Looks great!",
// 				DatePublished: time.Now(),
// 			},
// 		},
// 	}
// 	Tmpl.ExecuteTemplate(w, "post", p)
// }

// HandleNew handles preview logic
func HandleNew(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		Tmpl.ExecuteTemplate(w, "new", nil)
	case "POST":
		title := r.FormValue("title")
		content := r.FormValue("content")
		contentType := r.FormValue("content-type")
		r.ParseForm()
		if contentType == "page" {
			p := &Page{
				Title:   title,
				Content: content,
			}
			_, err := CreatePage(p)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			Tmpl.ExecuteTemplate(w, "page", p)
			return
		}
		// if contentType == "post" {
		// 	Tmpl.ExecuteTemplate(w, "post", &Post{
		// 		Title:   title,
		// 		Content: content,
		// 	})
		// 	return
		// }
	default:
		http.Error(w, "Method not supported:"+r.Method, http.StatusMethodNotAllowed)
	}
}