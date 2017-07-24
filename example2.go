package main

import (
	"context"
	"fmt"
	logfaceClient "github.com/Polarishq/logface-sdk-go/client"
	logfaceEvents "github.com/Polarishq/logface-sdk-go/client/events"
	logfaceModels "github.com/Polarishq/logface-sdk-go/models"
	rtclient "github.com/go-openapi/runtime/client"
	"time"
)

func fancy_sender(customContext context.Context) (chan logfaceModels.Event){
	inchan := make(chan logfaceModels.Event)
	authI := rtclient.BasicAuth(
		"Client ID",
		"Client Secret")
	httpClient := logfaceClient.NewHTTPClient(nil).Events

	var allEvents []*logfaceModels.Event
	params := logfaceEvents.NewEventsParams()
	params.Context = customContext
	go func() {
		for event := range inchan {
			allEvents = append(allEvents, &event)
			params.Events = allEvents
			eok, err := httpClient.Events(params, authI)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Printf("Success: %v\n", eok)
			}
		}
	}()
	return inchan
}

func main(){
	var customContext context.Context
	var cancel context.CancelFunc
	var event logfaceModels.Event
	customContext, cancel = context.WithTimeout(context.Background(), 4000*time.Millisecond)
	defer cancel()

	outchan := fancy_sender(customContext)
	for i := 1; i <= 10; i++ {
		event_payload := make(map[string]string)
		event_payload["raw"] = fmt.Sprintf("my event %v", i)
		event = logfaceModels.Event{
			Event: event_payload,
		}
		outchan <- event
	}
}
