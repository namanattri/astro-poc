package main

import (
	"astro-poc/utils"
	"fmt"

	"github.com/mshafiee/swephgo"
)

func PrintPlanetaryPositions() {
	year := 1989
	month := 5
	day := 21
	hour := 16.0667 // Afternoon
	// Birth location
	// lat := 30.9089 // Latitude of Solan
	// lon := 77.0953 // Longitude of Solan
	julianDay := swephgo.Julday(year, month, day, hour, swephgo.SeGregCal)
	fmt.Printf("date: %02d.%02d.%d at 0:00 Universal time\n", day, month, year)
	fmt.Printf("planet \tlongitude\tlongitude (DMS)\t\t\tlatitude\tlatitude (DMS)\tdistance\tspeed long.\n")

	for p := swephgo.SeSun; p <= swephgo.SePluto; p++ {
		if p == swephgo.SeEarth {
			continue
		}
		iflag := 0
		x2 := make([]float64, 6)
		serr := make([]byte, 256)
		iflgret := swephgo.CalcUt(julianDay, p, iflag, x2, serr)

		if iflgret < 0 {
			fmt.Printf("Error: %s\n", serr)
		}

		snam := make([]byte, 256)
		swephgo.GetPlanetName(p, snam)

		// Convert longitude and latitude to DMS format
		longitudeDMS := utils.ConvertToDMSString(x2[0], utils.Longitude)
		latitudeDMS := utils.ConvertToDMSString(x2[1], utils.Latitude)

		fmt.Printf("%10s\t%11.7f\t%s\t\t\t%10.7f\t%s\t%10.7f\t%10.7f\n", snam, x2[0], longitudeDMS, x2[1], latitudeDMS, x2[2], x2[3])
	}
}
