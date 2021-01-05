package models


type Direction int
const (
	North Direction = iota
	East
	South
	West
)

type GridIdentifier struct {
	Northings byte
	Eastings byte
}

type MeterGrid struct {
	Northings uint32 // 5 digits, meters
	Eastings uint32 // 5 digits, meters
}

type Zone struct {
	ZoneID GridIdentifier
	Sectors []Sector
	LengthNorth float64
	LengthEast float64
	SectorsNorth byte
	SectorsEast byte
}

type Sector struct {
	ZoneID GridIdentifier
	SectorID GridIdentifier
	Locations map[MeterGrid]*Location
}

type Location struct {
	ZoneID GridIdentifier // 6-degree Polar
	SectorID GridIdentifier // 100km square ID within Polar Zone ID on earth
	Meters MeterGrid
	Elevation float32
}


/*
Let's talk Cylinders
Coordinates with Location Struct...

if d < 100000 ( 100 km, which is rather excessive ):
	zone_size = 1000 ( 1 km )
if d > 100000:
	f = int(d/100000)
	zone_size = f * 1000

With Zones being a mere KM (or a few KM) Sectors are not *really* necessary
sector_size = zone_size // 1:1

Northings = meters within sector (Y)
Eastings = meters within sector (X)
 */