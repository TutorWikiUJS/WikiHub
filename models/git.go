package models

import (
	//"encoding/json"
	"fmt"
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

type Form struct {
	info  map[string]string
	index []string
}

func InitClient() *http.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: Token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	return tc
}

func CreateNewFile(form Form, img *os.File, tc *http.Client) (string, int) {

	name := form.info["name"]
	desc := fmt.Sprintf("#%s\n---\n\n", form.info["name"])

	for _, v := range form.index {
		if form.info[v] != "" && v != "detail" {
			desc = desc + fmt.Sprintf("- %s:%s\n", v, form.info[v])
		}
	}

	desc = desc + "\n---\n- 其他:" + form.info["detail"]

	log.Println("name:"+name, "detail:"+desc)
	client := github.NewClient(tc)

	commitMessage := "Add " + name

	message := github.RepositoryContentFileOptions{Message: &commitMessage, Content: []byte(desc)}

	uploadName := "tutors/" + name + strconv.Itoa(int(time.Now().Unix())) + ".md"

	_, r, err := client.Repositories.CreateFile(Name, Repo, uploadName, &message)
	if err != nil {
		log.Println(err)
		return "", -1
	} else {
		log.Println(r)
		return uploadName, 0
	}
}

func (this *Form) Add(key, value string) {
	if this.info == nil {
		this.info = make(map[string]string)
	}
	this.index = append(this.index, key)
	this.info[key] = value
}
