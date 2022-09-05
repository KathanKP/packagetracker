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
