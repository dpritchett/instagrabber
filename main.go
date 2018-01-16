package main

import (
	"./instagram"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := ":" + os.Getenv("PORT")
	http.HandleFunc("/", printPhotoUrl)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func printPhotoUrl(w http.ResponseWriter, r *http.Request) {
	photoUrl := instagram.GetInstaPhotoUrl()
	fmt.Fprintln(w, "<html><body><img src='"+photoUrl+"'/>")
}
