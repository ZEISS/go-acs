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

	"github.com/zeiss/go-acs/calls"
	internal "github.com/zeiss/go-acs/events"
	"github.com/zeiss/pkg/conv"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/zeiss/go-acs"
)

var (
	endpointURL string = ""
	key         string = ""
)

func main() {
	client := http.Client{}

	sys := make(chan os.Signal, 1)
	signal.Notify(sys, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	acsClient := acs.New(endpointURL, key, &client)

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
				case "Microsoft.Communication.RecognizeCompleted":
					event := &internal.MicrosoftCommunicationCallConnected{}
					err := e.DataAs(event)
					if err != nil {
						log.Fatalln(err)
						continue
					}

					err = acsClient.Call.CallHangUp(ctx, event.CallConnectionID)
					if err != nil {
						log.Fatalf("Error hanging up call: %v", err)
					}
				case "Microsoft.Communication.ParticipantsUpdated":
					event := &internal.MicrosoftCommunicationParticipantsUpdated{}
					err := e.DataAs(event)
					if err != nil {
						log.Fatalln(err)
						continue
					}

					for _, p := range event.Participants {
						if p.Identifier.Kind != "communicationUser" {
							continue
						}

						req := &calls.CallRecognizeRequest{
							RecognizeInputType: calls.RecognizeInputTypeChoices,
							PlayPrompt: calls.PlaySource{
								Kind: calls.PlaySourceTypeText,
								TextSource: &calls.TextSource{
									Text:      "Hello, the following incident occured: Instance-12345 on Azure is down. Please press 1 to acknowledge or 0 to decline.",
									VoiceName: "en-US-AriaNeural",
								},
							},
							RecognizeOptions: &calls.RecognizeOptions{
								InterruptPrompt:                true,
								InitialSilenceTimeoutInSeconds: 60,
								Choices: []calls.Choice{
									{
										Label:   "Acknowledged",
										Phrases: []string{"Acknowledge", "Ack", "Yes"},
										Tone:    calls.ToneOne,
									},
									{
										Label:   "Declined",
										Phrases: []string{"Decline", "No"},
										Tone:    calls.ToneZero,
									},
								},
								TargetParticipant: &calls.CommunicationIdentifier{
									Kind: conv.String(p.Identifier.Kind),
									CommunicationUser: &calls.CommunicationUser{
										ID: p.Identifier.CommunicationUser.ID,
									},
								},
							},
							OperationCallbackUri: "",
						}

						fmt.Println(req)

						err = acsClient.Call.CallMediaRecognize(ctx, event.CallConnectionID, key, req)
						if err != nil {
							panic(err)
						}
					}
				case "Microsoft.Communication.CallConnected":

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
			Value: "+",
		},
		CallIntelligenceOptions: &calls.CallIntelligenceOptions{
			CognitiveServicesEndpoint: "",
		},
		Targets: []calls.CommunicationIdentifier{
			{
				Kind: "phonenumber",
				PhoneNumber: &calls.PhonenumberIdentifier{
					ID:    "+",
					Value: "+",
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
