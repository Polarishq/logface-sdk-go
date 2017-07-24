package main

import (
	"fmt"
	logfaceClient "github.com/Polarishq/logface-sdk-go/client"
	logfaceClientEvents "github.com/Polarishq/logface-sdk-go/client/events"
	logfaceModels "github.com/Polarishq/logface-sdk-go/models"
	runtimeClient "github.com/go-openapi/runtime/client"
)

func main() {
	basicAuth := runtimeClient.BasicAuth(
		"Client ID",
		"Client Secret")

	singleEvent := logfaceModels.Event{
		Event:  map[string]string{"raw": "simple logface test"},
		Source: "my computer",
	}
	events := logfaceModels.Events{}
	events = append(events, &singleEvent)

	params := logfaceClientEvents.NewEventsParams()
	params.Events = events

	httpClient := logfaceClient.NewHTTPClient(nil).Events
	eok, err := httpClient.Events(params, basicAuth)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Success: %v\n", eok)
	}
}
