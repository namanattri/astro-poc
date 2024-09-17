package utils

import (
	"fmt"
	"math"
)

// CoordinateType defines whether the input is latitude or longitude
type CoordinateType string

const (
	Latitude  CoordinateType = "latitude"
	Longitude CoordinateType = "longitude"
)

// ConvertToDMS converts a decimal coordinate (latitude or longitude) to degrees, minutes, and seconds.
func ConvertToDMS(decimal float64, coordType CoordinateType) (int, int, float64, string) {
	// Adjust the input based on coordinate type
	switch coordType {
	case Latitude:
		// Ensure latitude stays within -90 and 90
		if decimal > 90 {
			decimal = 90
		} else if decimal < -90 {
			decimal = -90
		}
	case Longitude:
		// Wrap the longitude to the range of -180 to 180
		decimal = wrapLongitude(decimal)
	}

	// Determine direction based on the coordinate type and value
	var direction string
	if coordType == Latitude {
		if decimal >= 0 {
			direction = "N"
		} else {
			direction = "S"
		}
	} else if coordType == Longitude {
		if decimal >= 0 {
			direction = "E"
		} else {
			direction = "W"
		}
	}

	// Make decimal positive for calculation
	decimal = math.Abs(decimal)

	// Get the degrees by truncating the decimal part
	degrees := int(decimal)

	// Calculate the remaining decimal part and convert to minutes
	minutesFloat := (decimal - float64(degrees)) * 60
	minutes := int(minutesFloat)

	// Calculate the remaining decimal part and convert to seconds
	seconds := (minutesFloat - float64(minutes)) * 60

	return degrees, minutes, seconds, direction
}

// ConvertToDMSString converts a decimal coordinate to a formatted DMS string.
func ConvertToDMSString(decimal float64, coordType CoordinateType) string {
	// Get the DMS values
	degrees, minutes, seconds, direction := ConvertToDMS(decimal, coordType)

	// Format the DMS into a string
	return fmt.Sprintf("%d° %d' %.2f\" %s", degrees, minutes, seconds, direction)
}

// wrapLongitude ensures that longitude values are within the range of -180 to 180 degrees
func wrapLongitude(longitude float64) float64 {
	for longitude > 180 {
		longitude -= 360
	}
	for longitude < -180 {
		longitude += 360
	}
	return longitude
}
