package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"

	apiai "github.com/mlabouardy/dialogflow-go-client/models"
	. "github.com/mlabouardy/dialogflow-watchnow-messenger/models"
	. "github.com/mlabouardy/moviedb"
)

const (
	FACEBOOK_API   = "https://graph.facebook.com/v2.6/me/messages?access_token=%s"
	IMAGE          = "http://37.media.tumblr.com/e705e901302b5925ffb2bcf3cacb5bcd/tumblr_n6vxziSQD11slv6upo3_500.gif"
	VISIT_SHOW_URL = "http://labouardy.com"
)

func BuildCarousel(shows []Show) []Element {
	elements := make([]Element, 0)
	for _, show := range shows {
		element := Element{
			Title:    show.Title,
			ImageURL: show.Cover,
			DefaultAction: DefaultAction{
				Type: "web_url",
				URL:  VISIT_SHOW_URL,
			},
			Buttons: []Button{
				Button{
					Type:  "web_url",
					Title: "Watch now",
					URL:   VISIT_SHOW_URL,
				},
			},
		}
		elements = append(elements, element)
	}
	return elements
}

func ProcessMessage(event Messaging) {
	var userQuery = event.Message.Text
	var dialogFlowResponse = GetResponse(userQuery)
	client := &http.Client{}
	var response Response

	if !reflect.DeepEqual(dialogFlowResponse.Metadata, apiai.Metadata{}) && dialogFlowResponse.Metadata.IntentName == "shows" {
		var showType = dialogFlowResponse.Parameters["show-type"]
		db := NewMovieDB()
		var shows []Show

		if showType == "movie" {
			shows = db.GetNowPlayingMovies()
		} else {
			shows = db.GetAiringTodayShows()
		}

		response = Response{
			Recipient: User{
				ID: event.Sender.ID,
			},
			Message: Message{
				Attachment: &Attachment{
					Type: "template",
					Payload: Payload{
						TemplateType: "generic",
						Elements:     BuildCarousel(shows[:10]),
					},
				},
			},
		}
	} else {
		response = Response{
			Recipient: User{
				ID: event.Sender.ID,
			},
			Message: Message{
				Text: dialogFlowResponse.Fulfillment.Speech,
			},
		}
	}

	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(&response)

	url := fmt.Sprintf(FACEBOOK_API, os.Getenv("PAGE_ACCESS_TOKEN"))
	req, err := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}
