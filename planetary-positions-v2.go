package main

import (
	"fmt"
	"log"

	"github.com/mshafiee/swephgo"
)

func PrintPlanetaryPositionsV2() {
	// Initialize Swiss Ephemeris
	swephgo.SetEphePath([]byte("ephe/")) // set path to ephemeris files, empty string uses default

	// Input birth details
	year := 1989
	month := 5
	day := 21
	hour := 16.0667 // 12 PM, can include fractional part for minutes/seconds

	// Birth location
	// lat := 30.9089 // Latitude of Solan
	// lon := 77.0953 // Longitude of Solan

	// Convert to Julian Day Number
	jd := swephgo.Julday(year, month, day, hour, swephgo.SeGregCal)

	// Flags for calculations
	planetFlags := swephgo.SeflgSwieph | swephgo.SeTrueNode // Swiss Ephemeris with true node

	// Planets positions (Sun, Moon, Mercury, Venus, Mars, etc.)
	xx := make([]float64, 6)
	serr := make([]byte, 256)
	fmt.Printf("planet \tPosition\tlongitude\tlatitude\tdistance\tspeed long.\n")
	for planet := swephgo.SeSun; planet <= swephgo.SePluto; planet++ {
		swephgo.Calc(jd, planet, planetFlags, xx, serr)
		if serr != nil {
			log.Fatalf("Error calculating planet: %v", serr)
		}
		fmt.Printf("%s\t%11.7f\t%10.7f\t%10.7f\t%10.7f\n", planetName(planet), xx[0], xx[1], xx[2], xx[2])
	}

	// // Calculate the Ascendant
	// ascendant, err := swephgo.CalcAsc(jd, lon, lat, 0)
	// if err != nil {
	// 	log.Fatalf("Error calculating Ascendant: %v", err)
	// }
	// fmt.Printf("Ascendant: %f degrees\n", ascendant)

	// // Calculate Lunar Nodes (True Node)
	// nodePos, err := swephgo.Calc(jd, swephgo.SE_TRUE_NODE, planetFlags)
	// if err != nil {
	// 	log.Fatalf("Error calculating lunar node: %v", err)
	// }
	// fmt.Printf("True Lunar Node: %f degrees\n", nodePos[0])

	// // Calculate Mean Lunar Node
	// meanNodePos, err := swephgo.Calc(jd, swephgo.SE_MEAN_NODE, planetFlags)
	// if err != nil {
	// 	log.Fatalf("Error calculating mean lunar node: %v", err)
	// }
	// fmt.Printf("Mean Lunar Node: %f degrees\n", meanNodePos[0])
}

// Helper function to get planet name
func planetName(planet int) string {
	switch planet {
	case swephgo.SeSun:
		return "Sun"
	case swephgo.SeMoon:
		return "Moon"
	case swephgo.SeMercury:
		return "Mercury"
	case swephgo.SeVenus:
		return "Venus"
	case swephgo.SeMars:
		return "Mars"
	case swephgo.SeJupiter:
		return "Jupiter"
	case swephgo.SeSaturn:
		return "Saturn"
	case swephgo.SeUranus:
		return "Uranus"
	case swephgo.SeNeptune:
		return "Neptune"
	case swephgo.SePluto:
		return "Pluto"
	default:
		return "Unknown"
	}
}
