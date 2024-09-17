package main

import (
	"astro-poc/lib"
	"fmt"

	"github.com/mshafiee/swephgo"
)

func main() {
	sweVer := make([]byte, 12)
	swephgo.Version(sweVer)
	swephgo.SetEphePath([]byte("./ephe"))
	fmt.Printf("Library used: Swiss Ephemeris v%s\n", sweVer)

	year := 1989
	month := 5
	day := 21
	hour := 16.0667 // Afternoon
	// Birth location
	lat := 30.9089 // Latitude of Solan
	lon := 77.0953 // Longitude of Solan
	julianDay := swephgo.Julday(year, month, day, hour, swephgo.SeGregCal)
	fmt.Printf("date: %02d.%02d.%d at 0:00 Universal time\n", day, month, year)
	lib.PrintPlanetaryPositions(julianDay)
	lib.PrintAscendant(lat, lon, julianDay)
}
