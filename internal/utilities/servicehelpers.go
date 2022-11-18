package utilities

import (
	"github.com/KathanKP/packagetracker/internal/services/fedex"
)

func GetLatestPackageStatus(trackingId string, serviceName string) string {
	return fedex.GetFedexPackageStatus(trackingId)
}
