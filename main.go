package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		newDomain := os.Getenv("NEW_DOMAIN")
		if newDomain == "" {
			newDomain = "https://google.com"
		} else {
			newDomain = os.Getenv("NEW_DOMAIN")
		}
		newURL := newDomain + r.URL.Path
		http.Redirect(w, r, newURL, http.StatusMovedPermanently)
	})
	http.ListenAndServe(":8080", nil)
}
