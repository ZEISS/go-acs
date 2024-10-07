package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/zeiss/go-acs"
	"github.com/zeiss/go-acs/sms"
)

var (
	endpointURL string = "https://acs-nova-demo.germany.communication.azure.com/"
	key         string
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
				To: "+4915120756627",
			},
		},
		Message: "Hello, world!",
		SMSSendOptions: sms.SMSSendOptions{
			EnableDeliveryReport: true,
			Tag:                  "example",
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
