package models

import (
	//"encoding/json"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	Token = "6fdd62cc432c2c3fc8f65f1a296733dfec5d3f71"
	Name  = "publicidujs"
	Repo  = "TutorWiki"
)

func InitClient() *http.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: Token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	return tc
}

func CreateNewFile(name, detail string, img *os.File, tc *http.Client) int {
	log.Println("name:"+name, "detail:"+detail)
	client := github.NewClient(tc)

	if name == "" || detail == "" {
		log.Println("Invaild Input")
		return -2
	}

	commitMessage := "Add " + name

	message := github.RepositoryContentFileOptions{Message: &commitMessage, Content: []byte(detail)}

	uploadName := "tutors/" + name + strconv.Itoa(int(time.Now().Unix())) + ".md"

	_, r, err := client.Repositories.CreateFile(Name, Repo, uploadName, &message)
	if err != nil {
		log.Println(err)
		return -1
	} else {
		log.Println(r)
		return 0
	}
}
