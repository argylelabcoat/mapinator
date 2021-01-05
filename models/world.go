package models

import (
	"errors"
	"math"
)

const (
	KM float64 = 1000
	KM100 float64 = 100000
)

type World struct {
	Terrain map[*Location]Terrain
	Roads map[*Location]Road
	Zones map[GridIdentifier]Zone
}

// Sectors are 100KM squares
// Equator rows starts with letter A in odd zones, F in even zones
// At Equator, columns are A-H within Zone 1, J-R (w/o O) in Zone 2, and S-Z in Zone 3, then repeat
// Letters I and O are excluded

// Zones are defined via Lat/Lon Lines, 6 degree / 8 degree on Earth



func Cylinder(diameter float64, length float64) (*World, error) {
	var world World
	var circumference float64

	if diameter < KM {
		return nil, errors.New("insufficient diameter")
	}
	if length < KM {
		return nil, errors.New("insufficient length")
	}
	circumference = diameter * math.Pi
	world.SurfaceArea = circumference * length

	if circumference > KM100 {
		// We need Sectors
		if circumference > 6 * KM100 {
			// We need Zones!
		} else {
			// Just Sectors
		}
	} else {
		// No Sectors E/W
	}

	if length > KM100 {
		// We need Sectors
		if length > 6 * KM100 {
			// We need Zones!
		} else {
			// Just Sectors
		}
	} else {
		// No Sectors N/S
	}


	return &world, nil
}