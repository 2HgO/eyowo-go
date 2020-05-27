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

	var unknownUser = ""
	res, err := client.ValidateUser(unknownUser)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res)
}
