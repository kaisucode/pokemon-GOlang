package maploader

import (
	"io/ioutil"
	"kaisu/pokemon/constants"

	"kaisu/pokemon/src/mapparser"
	. "kaisu/pokemon/src/utils"

	tl "github.com/JoelOtter/termloop"
)

func LoadMapLevel(fileurl string) {
	// constants.CONSOLE_TEXT.SetText("loading map level of file: ")

	// create a base level
	constants.CURLEVEL = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
		Ch: '.',
	})

	lmap, err := ioutil.ReadFile(fileurl)
	CheckErr(err)

	parsers := make(map[string]tl.EntityParser)
	parsers["Player"] = mapparser.ParsePlayer
	parsers["Building"] = mapparser.ParseBuildings // handles collision rects and warp points
	// parsers["Npc"] = mapparser.ParsePlayer
	// parsers["Grass"] = mapparser.ParsePlayer
	err = tl.LoadLevelFromMap(string(lmap), parsers, constants.CURLEVEL)
	CheckErr(err)

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
	parsers["Player"] = mapparser.ParsePlayer
	parsers["Room"] = mapparser.ParseRooms
	// parsers["Npc"] = mapparser.ParsePlayer
	// parsers["Grass"] = mapparser.ParsePlayer
	err = tl.LoadLevelFromMap(string(lmap), parsers, constants.CURLEVEL)
	CheckErr(err)
}
