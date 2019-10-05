package main

import (
	"net/http"

	"github.com/meoconbatu/cms"
	"github.com/meoconbatu/cmsgrpc/view"
)

func main() {
	http.HandleFunc("/new", view.HandleNew)
	http.HandleFunc("/page/", cms.ServePage)
	http.ListenAndServe(":3000", nil)
}
