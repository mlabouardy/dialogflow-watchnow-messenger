package main

import (
	"log"
	"os"

	. "github.com/mlabouardy/apiai-go-client"
	apiai "github.com/mlabouardy/apiai-go-client/models"
)

func GetResponse(input string) apiai.Result {
	err, client := NewApiAiClient(apiai.Options{
		AccessToken: os.Getenv("DIALOG_FLOW_TOKEN"),
	})
	if err != nil {
		log.Fatal(err)
	}

	query := apiai.Query{
		Query: input,
	}
	resp, err := client.QueryFindRequest(query)
	if err != nil {
		log.Fatal(err)
	}
	return resp.Result
}
