package main

import (
	"fmt"
	"github.com/dpritchett/instagrabber/instagram"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator

	addr := ":" + os.Getenv("PORT")
	http.HandleFunc("/", renderPhoto)
	http.HandleFunc("/raw", printRawPhotoUrl)

	log.Println("Serving on port " + os.Getenv("PORT") + "...")
	log.Fatal(http.ListenAndServe(addr, nil))
}

func renderPhoto(w http.ResponseWriter, r *http.Request) {
	photoUrl := getRandomPhoto()
	fmt.Fprintln(w, "<html><body><img src='"+photoUrl+"'/>")
}

func printRawPhotoUrl(w http.ResponseWriter, r *http.Request) {
	photoUrl := getRandomPhoto()
	fmt.Fprintln(w, photoUrl)
}

func getRandomPhoto() string {
	photos := instagram.GetPhotos()

	photo := photos[rand.Intn(len(photos))]

	return photo.Thumbnails[2].Source
}
