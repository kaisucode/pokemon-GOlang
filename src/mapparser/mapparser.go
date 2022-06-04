package mapparser

import (
	"kaisu/pokemon/constants/mapstate"
	"kaisu/pokemon/src/building"
	"kaisu/pokemon/src/room"
	. "kaisu/pokemon/src/utils"

	tl "github.com/JoelOtter/termloop"
)

type ParseObj struct {
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
	return &ParseObj{e}
}

func ParseRooms(data map[string]interface{}) tl.Drawable {

	// render buildings (should be done in building init)
	// pokemonCenter := building.InitBuilding(constants.CURLEVEL)
	// pokemonCenter := building.NewBuilding()

	newRoom := room.NewRoom(data)
	name := string(data["name"].(string))
	Use(newRoom, name)
	// mapstate.LOCATIONS[name] = newRoom
	// return newBuilding

	// generate warp point
	// return warp point

	e := tl.NewEntity(
		int(data["x"].(float64)),
		int(data["y"].(float64)),
		1, 1,
	)
	return &ParseObj{e}

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

func ParseBuildings(data map[string]interface{}) tl.Drawable {

	// render buildings (should be done in building init)
	// pokemonCenter := building.InitBuilding(constants.CURLEVEL)
	// pokemonCenter := building.NewBuilding()

	xPos := int(data["x"].(float64))
	yPos := int(data["y"].(float64))
	url := string(data["url"].(string))
	name := string(data["name"].(string))

	newBuilding := building.NewBuilding(data)
	mapstate.LOCATIONS[name] = newBuilding
	// return newBuilding

	// generate warp point
	// return warp point

	// warpPoint := building.NewWarppoint(xPos, yPos, url)
	warpPoint := building.NewWarppoint(xPos-1, yPos+3, url)
	Use(warpPoint, xPos, yPos)
	Use(xPos, yPos, url)

	// return &warpPoint

	e := tl.NewEntity(
		int(data["x"].(float64)),
		int(data["y"].(float64)),
		1, 1,
	)
	return &ParseObj{e}

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
