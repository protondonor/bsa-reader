package maps

var locationTypes = [15]string{
	"City", "Town", "Village", "Farmstead",
	"Labyrinth", "Temple", "Tavern", "Fortress",
	"Manor", "Shrine", "Ruins", "Shack",
	"Graveyard", "Coven", "Your Ship",
}

func (l LatitudeType) GetType() string {
	return locationTypes[l.Type]
}
