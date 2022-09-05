package fedex

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	test_url         = "https://apis-sandbox.fedex.com"
	api_key          = ""
	shipping_account = ""
	secret_key       = ""
	oauth_endpoint   = "oauth/token"
)

func getOAuthToken() {
	oauthRequestBody, err := json.Marshal(map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     api_key,
		"client_secret": secret_key,
	})
	if err != nil {
		log.Fatalln(err)
	}
	oauthRespone, err := http.Post(test_url+"/"+oauth_endpoint, "application/x-www-form-urlencoded", bytes.NewBuffer(oauthRequestBody))
	if err != nil {
		log.Fatalln(err)
	}
	defer oauthRespone.Body.Close()
	oauthResponseBody, err := ioutil.ReadAll(oauthRespone.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(oauthResponseBody))
}

func GetPackageStatus(trackingId string) {
	log.Println("Getting the oauth token")
	getOAuthToken()
}
