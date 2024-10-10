package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/zeiss/go-acs"
	"github.com/zeiss/go-acs/calls"
)

var (
	endpointURL string = ""
	token       string = ""
	key         string = ""
)

func main() {
	client := http.Client{}

	sys := make(chan os.Signal, 1)
	signal.Notify(sys, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	acsClient := acs.New(endpointURL, token, &client)

	go func() {
		http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
			b, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Printf("Body reading error: %v", err)
				return
			}
			defer r.Body.Close()

			var events []cloudevents.Event
			err = json.Unmarshal(b, &events)
			if err != nil {
				log.Fatalln(err)
				return
			}

			fmt.Println(events)

			for _, e := range events {
				switch e.Type() {
				case "Microsoft.Communication.CallConnected":
					event := &calls.CallConnectedData{}
					err := e.DataAs(event)
					if err != nil {
						log.Fatalln(err)
						continue
					}

					req := &calls.CallMediaPlayRequest{
						PlaySources: []calls.PlaySource{
							{
								Kind: calls.PlaySourceTypeText,
								TextSource: &calls.TextSource{
									Text:      "Hello, this is a test message.",
									VoiceName: "en-US-NancyNeural",
								},
							},
						},
						PlayOptions: &calls.PlayOptions{
							Loop: false,
						},
						OperationCallbackUri: "",
					}

					fmt.Println(req)

					err = acsClient.Call.CallMediaPlay(ctx, event.CallConnectionID, key, req)
					if err != nil {
						panic(err)
					}
				}
			}
		})

		http.ListenAndServe(":8080", nil)
	}()

	go func() {
		oscall := <-sys
		log.Printf("system call:%+v", oscall)
		cancel()
	}()

	req := &calls.CreateCallRequest{
		SourceCallerIdNumber: &calls.PhonenumberIdentifier{
			Value: "+18662311561",
		},
		CallIntelligenceOptions: &calls.CallIntelligenceOptions{
			CognitiveServicesEndpoint: "",
		},
		Targets: []calls.CommunicationIdentifier{
			{
				Kind: "phonenumber",
				PhoneNumber: &calls.PhonenumberIdentifier{
					ID:    "+18662311561",
					Value: "+18662311561",
				},
			},
		},
		CallbackUri: "",
	}

	err := acsClient.Call.CreateCall(ctx, key, req)
	if err != nil {
		panic(err)
	}

	<-ctx.Done()
}
