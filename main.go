package main

import (
	"fmt"

	"github.com/mshafiee/swephgo"
)

func main() {
	sweVer := make([]byte, 12)
	swephgo.Version(sweVer)
	swephgo.SetEphePath([]byte("./ephe"))
	fmt.Printf("Library used: Swiss Ephemeris v%s\n", sweVer)

	printPlantPositions()
}

func printPlantPositions() {
	year := 2024
	month := 7
	day := 8
	hour := 12.0 // Noon
	julianDay := swephgo.Julday(year, month, day, hour, swephgo.SeGregCal)
	fmt.Printf("date: %02d.%02d.%d at 0:00 Universal time\n", day, month, year)
	fmt.Printf("planet \tlongitude\tlatitude\tdistance\tspeed long.\n")

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

		fmt.Printf("%10s\t%11.7f\t%10.7f\t%10.7f\t%10.7f\n", snam, x2[0], x2[1], x2[2], x2[3])
	}
}
