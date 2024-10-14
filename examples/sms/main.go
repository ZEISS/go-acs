package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/zeiss/go-acs"
	"github.com/zeiss/go-acs/sms"
)

var (
	endpointURL string = ""
	key         string = ""
)

func main() {
	client := http.Client{}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	acsClient := acs.New(endpointURL, key, &client)

	res, err := acsClient.SMS.SendSMS(ctx, &sms.Request{
		From: "+1234567890",
		SMSRecipients: []sms.SMSRecipients{
			{
				To: "+10987654321",
			},
		},
		Message: "Thanks for using our service!",
		SMSSendOptions: sms.SMSSendOptions{
			EnableDeliveryReport: true,
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
