package models

type RoadType int
const (
	PedestrianPath RoadType = iota
	BikePath
	ZeroGTrain
	Railway
)

type Road struct {
	Direction Direction
	Type RoadType
	Location *Location
}
