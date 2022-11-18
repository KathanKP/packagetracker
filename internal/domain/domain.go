package domain

import (
	"time"
)

type Package struct {
	Service       string
	TrackingId    string
	CurrentStatus string
	LastUpdated   time.Time
}

type FedexAuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

type FedexTrackingNumberInfo struct {
	TrackingNumber string `json:"trackingNumber"`
}

type FedexTrackingInfoPackage struct {
	TrackingNumberInfo FedexTrackingNumberInfo `json:"trackingNumberInfo"`
}

type FedexTrackingRequest struct {
	IncludeDetailedScans bool                       `json:"includeDetailedScans"`
	TrackingInfo         []FedexTrackingInfoPackage `json:"trackingInfo"`
}
