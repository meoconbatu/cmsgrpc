package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/meoconbatu/cmsgrpc/view"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/new", view.HandleNew)
	http.HandleFunc("/page/", view.ServePage)
	http.HandleFunc("/login", view.ServeLogin)
	http.HandleFunc("/register", view.ServeRegister)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
