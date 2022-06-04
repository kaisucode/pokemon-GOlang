package mapparser

import (
	"io/ioutil"
	"kaisu/pokemon/constants"
	"kaisu/pokemon/constants/mapstate"
	"kaisu/pokemon/constants/playerstate"
	"kaisu/pokemon/src/building"
	"kaisu/pokemon/src/player"
	"kaisu/pokemon/src/room"
	. "kaisu/pokemon/src/utils"

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

func LoadMapLevel(fileurl string) {

	// create a base level
	constants.CURLEVEL = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
		Ch: '.',
	})

	lmap, err := ioutil.ReadFile(fileurl)
	CheckErr(err)

	parsers := make(map[string]tl.EntityParser)
	parsers["Player"] = ParsePlayer
	parsers["Building"] = ParseBuildings // handles collision rects and warp points
	// parsers["Npc"] = mapparser.ParsePlayer
	// parsers["Grass"] = mapparser.ParsePlayer
	err = tl.LoadLevelFromMap(string(lmap), parsers, constants.CURLEVEL)
	CheckErr(err)

	// render player
	playerstate.PLAYER = &player.Player{
		Entity: tl.NewEntity(1, 1, 1, 1),
	}

	// Set the character at position (0, 0) on the entity.
	playerstate.PLAYER.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '@'})
	constants.CURLEVEL.AddEntity(playerstate.PLAYER)
}

func LoadRoom(fileurl string) {

	// create a base level
	constants.CURLEVEL = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorGreen,
	})

	lmap, err := ioutil.ReadFile(fileurl)
	CheckErr(err)

	parsers := make(map[string]tl.EntityParser)
	parsers["Player"] = ParsePlayer
	parsers["Room"] = ParseRooms
	// parsers["Npc"] = mapparser.ParsePlayer
	// parsers["Grass"] = mapparser.ParsePlayer
	err = tl.LoadLevelFromMap(string(lmap), parsers, constants.CURLEVEL)
	CheckErr(err)

	// render player
	playerstate.PLAYER = &player.Player{
		Entity: tl.NewEntity(1, 1, 1, 1),
	}

	// Set the character at position (0, 0) on the entity.
	playerstate.PLAYER.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '@'})
	constants.CURLEVEL.AddEntity(playerstate.PLAYER)
}
