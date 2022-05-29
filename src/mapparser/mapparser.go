package mapparser

import (
	"kaisu/pokemon/constants/mapstate"
	"kaisu/pokemon/src/building"

	tl "github.com/JoelOtter/termloop"
)

type Player struct {
	*tl.Entity
}

// Here we define a parse function for reading a Player out of JSON.
func ParsePlayer(data map[string]interface{}) tl.Drawable {
	e := tl.NewEntity(
		int(data["x"].(float64)),
		int(data["y"].(float64)),
		1, 1,
	)
	e.SetCell(0, 0, &tl.Cell{
		Ch: []rune(data["ch"].(string))[0],
		Fg: tl.Attr(data["color"].(float64)),
	})
	return &Player{e}
}

func ParseBuildings(data map[string]interface{}) tl.Drawable {

	// render buildings (should be done in building init)
	// pokemonCenter := building.InitBuilding(constants.CURLEVEL)
	// pokemonCenter := building.NewBuilding()

	newBuilding := building.NewBuilding(data)
	name := string(data["name"].(string))

	mapstate.LOCATIONS[name] = newBuilding
	// return newBuilding

	e := tl.NewEntity(
		int(data["x"].(float64)),
		int(data["y"].(float64)),
		1, 1,
	)
	return &Player{e}

	// e := tl.NewEntity(
	//   int(data["x"].(float64)),
	//   int(data["y"].(float64)),
	//   1, 1,
	// )
	// e.SetCell(0, 0, &tl.Cell{
	//   Ch: []rune(data["ch"].(string))[0],
	//   Fg: tl.Attr(data["color"].(float64)),
	// })
	// return &Player{e}
}
