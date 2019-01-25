package service

import (
	"encoding/json"
	"fmt"
	"goCanvas/pkg/instructure/models"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Provider struct {

	BaseURL		string
	AccessToken	string

}

func (p *Provider) formatUrl(url string) string {

	if !strings.HasPrefix(p.BaseURL, "https") {
		return "https://" + url
	}

	return url
}

func (p *Provider) FetchProfile() (*models.Profile, bool) {

	url := p.formatUrl(p.BaseURL + "users/self/profile?access_token=" +
		p.AccessToken)

	response, err := http.Get(fmt.Sprint(url))

	if err != nil {
		fmt.Println(err.Error())
		return nil, false
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
		return nil, false
	}

	var profile models.Profile
	err = json.Unmarshal([]byte(responseData), &profile)

	if err != nil {
		return nil, false
	}

	return &profile, true
}

func (p *Provider) FetchCourses() (*[]models.Course, bool) {

	url := p.formatUrl(p.BaseURL + "courses?access_token=" +
		p.AccessToken)

	response, err := http.Get(fmt.Sprint(url))

	if err != nil {
		fmt.Println(err.Error())
		return nil, false
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
		return nil, false
	}

	var courses []models.Course
	err = json.Unmarshal([]byte(responseData), &courses)

	if err != nil {
		return nil, false
	}

	return &courses, true
}
