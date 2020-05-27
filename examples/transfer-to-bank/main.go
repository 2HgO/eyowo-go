package main

import (
	"log"

	"github.com/2HgO/eyowo-go"
)

func main() {
	var (
		appKey                  = ""
		appSecret               = ""
		mobile                  = ""
		cachedRefreshToken      = ""
		transferAmount     uint = 100000 // #1000.00 in kobo
		recipientName           = "gogo"
		recipientAccount        = "001324512"
		bankCode                = "000013"
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

	res, err := client.TransferToBank(transferAmount, recipientName, recipientAccount, bankCode)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res)
}
