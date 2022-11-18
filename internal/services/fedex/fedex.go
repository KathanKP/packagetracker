package fedex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/KathanKP/packagetracker/internal/domain"
)

const (
	test_url         = "https://apis-sandbox.fedex.com"
	api_key          = "l7113606b2c820472ea61623b456de702e"
	shipping_account = "740561073"
	secret_key       = "fa4b21423f484e7e8e310b7e386a9a47"
	oauth_endpoint   = "oauth/token"
	track_endpoint   = "track/v1/trackingnumbers"
)

func getOAuthToken() (string, error) {
	oauthRequestBody := url.Values{}
	oauthResponseBody := domain.FedexAuthResponse{}

	oauthRequestBody.Set("grant_type", "client_credentials")
	oauthRequestBody.Set("client_id", api_key)
	oauthRequestBody.Set("client_secret", secret_key)
	encodedRequestBody := oauthRequestBody.Encode()

	oauthRespone, err := http.Post(test_url+"/"+oauth_endpoint, "application/x-www-form-urlencoded", strings.NewReader(encodedRequestBody))
	if err != nil {
		log.Fatalln(err)
	}
	defer oauthRespone.Body.Close()
	err = json.NewDecoder(oauthRespone.Body).Decode(&oauthResponseBody)
	if err != nil {
		return "", err
	}
	return oauthResponseBody.AccessToken, nil
}

func GetFedexPackageStatus(trackingId string) string {
	log.Println("Getting the oauth token")
	oauthToken, err := getOAuthToken()
	//Get and return the package status
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Fedex oauth token is: " + oauthToken)
	trackingObject := domain.FedexTrackingRequest{
		IncludeDetailedScans: false,
		TrackingInfo: []domain.FedexTrackingInfoPackage{
			{
				TrackingNumberInfo: domain.FedexTrackingNumberInfo{TrackingNumber: trackingId},
			},
		},
	}
	fmt.Printf("Getting tracking details for these packages %+v.\n", trackingObject)
	jsonRequestObject, err := json.Marshal(trackingObject)
	if err != nil {
		log.Fatalln(err)
	}

	// headers
	client := &http.Client{}
	req, err := http.NewRequest("POST", test_url+"/"+track_endpoint, bytes.NewBuffer(jsonRequestObject))
	if err != nil {
		log.Fatalln(err)
	}
	// var bearerToken := "Bearer " +
	req.Header = http.Header{
		"x-customer-transaction-id": {"123467"},
		"Content-type":              {"application/json"},
		"X-locale":                  {"en_US"},
		"Authorization":             {"Bearer " + oauthToken},
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println(string([]byte(body)))
	return ""
}
