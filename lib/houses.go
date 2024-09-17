package lib

import (
	"astro-poc/utils"
	"fmt"

	"github.com/mshafiee/swephgo"
)

func PrintAscendant(geolat, geolon, julianDay float64) {
	// Time variables
	hsys := 'P' // Placidus house system (can also use 'K' for Koch, 'R' for Regiomontanus, etc.)

	// Create slice for cusp positions and ascmc (ascendant, midheaven, etc.)
	cusps := make([]float64, 13) // 12 houses + 1 (cusps[0] is unused)
	ascmc := make([]float64, 10) // Contains the ascendant and other calculated points

	// Call the Swiss Ephemeris function for house calculation
	ret := swephgo.Houses(julianDay, geolat, geolon, int(hsys), cusps, ascmc)

	// Check if the function returned successfully
	if int32(ret) == 0 {
		// Print the ASC (ascendant) which is at index 0 of ascmc
		fmt.Printf("Ascendant (ASC) is: %.6f\t%s\n", ascmc[0], utils.ConvertToDMSString(ascmc[0], utils.Longitude))
	} else {
		fmt.Println("Error calculating ascendant.")
	}
}
