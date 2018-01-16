package main

import (
	"./instagram"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator

	addr := ":" + os.Getenv("PORT")
	http.HandleFunc("/", printPhotoUrl)

	log.Println("Serving on port " + os.Getenv("PORT") + "...")
	log.Fatal(http.ListenAndServe(addr, nil))
}

func printPhotoUrl(w http.ResponseWriter, r *http.Request) {
	photos := instagram.GetPhotos()

	photo := photos[rand.Intn(len(photos))]

	photoUrl := photo.Thumbnails[2].Source

	fmt.Fprintln(w, "<html><body><img src='"+photoUrl+"'/>")
}
