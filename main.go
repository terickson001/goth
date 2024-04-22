package main

import (
	"log"
	"net/http"

	"terickson001/painterly/view/partial"
	"terickson001/painterly/view"
	"terickson001/painterly/view/layout"

	"github.com/a-h/templ"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	c := layout.Base(view.Index())
	http.Handle("/", templ.Handler(c))

	http.Handle("/foo", templ.Handler(partial.Foo()))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
