package mapstate

import (
	"kaisu/pokemon/src/building"
)

type BuildingEntries = map[string](*building.Building)

var LOCATIONS BuildingEntries
