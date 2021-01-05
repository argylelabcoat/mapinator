package models

type TerrainComposition int
const (
	Earth TerrainComposition = iota
	Water
	Pavement
)

type Terrain struct {
	Elevation float32
	Composition TerrainComposition
	Location *Location
}
