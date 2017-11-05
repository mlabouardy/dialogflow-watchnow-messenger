package main

import (
	"log"
	"os"

	. "github.com/mlabouardy/dialogflow-go-client"
	apiai "github.com/mlabouardy/dialogflow-go-client/models"
)

func GetResponse(input string) apiai.Result {
	err, client := NewDialogFlowClient(apiai.Options{
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
