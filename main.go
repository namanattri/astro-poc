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

	PrintPlanetaryPositions()
}
