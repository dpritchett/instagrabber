package instagram

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Thumbnail struct {
	Source string `json:"src"`
}

type InstagramNode struct {
	DisplaySrc string      `json:"display_src"`
	Thumbnails []Thumbnail `json:"thumbnail_resources"`
}

type InstagramMedia struct {
	Nodes []InstagramNode `json:"nodes"`
}

type InstagramFeed struct {
	User struct {
		FullName string         `json:"full_name"`
		Media    InstagramMedia `json:"media"`
	}
}

func GetInstaPhotoUrl() string {
	photos := GetPhotos()

	thumbs := photos[0].Thumbnails

	photo := thumbs[2]

	log.Printf("+%v", photo)

	photoUrl := photo.Source

	log.Println(photoUrl)

	return photoUrl
}

func GetPhotos() []InstagramNode {
	feedUrl := "https://www.instagram.com/danieljpritchett/?__a=1"
	resp, err := http.Get(feedUrl)

	if err != nil {
		log.Println("Unable to fetch instagram feed: ")
		log.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("Unable to read instagram response:")
		log.Println(err)
	}

	var instaFeed InstagramFeed

	err = json.Unmarshal(body, &instaFeed)

	if err != nil {
		log.Println("Unable to parse instafeed:")
		log.Println(err)
	}

	return instaFeed.User.Media.Nodes
}
