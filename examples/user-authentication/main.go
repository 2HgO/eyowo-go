package main

import (
	"log"

	"github.com/2HgO/eyowo-go"
)

func main() {
	var (
		appKey    = ""
		appSecret = ""
		mobile    = ""
	)
	client, err := eyowo.NewClient(appKey, appSecret, mobile, eyowo.PRODUCTION)
	if err != nil {
		log.Fatalln(err)
	}

	// Start authentication flow
	res, err := client.AuthenticateUser("sms")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res)

	// Use passcode received to get access and refresh tokens
	res, err = client.AuthenticateUser("sms", "passcode")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res)
	log.Println(client.GetAccessToken())
	log.Println(client.GetRefreshToken())
}
