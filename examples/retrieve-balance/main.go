package main

import (
	"log"

	"github.com/2HgO/eyowo-go"
)

func main() {
	var (
		appKey             = ""
		appSecret          = ""
		mobile             = ""
		cachedRefreshToken = ""
	)
	client, err := eyowo.NewClient(appKey, appSecret, mobile, eyowo.PRODUCTION)
	if err != nil {
		log.Fatalln(err)
	}

	client.SetRefreshToken(cachedRefreshToken)
	err = client.RefreshAccessToken()
	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.GetBalance()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res)
}
