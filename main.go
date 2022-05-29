package main

import (
	"io/ioutil"
	"kaisu/pokemon/constants"
	"kaisu/pokemon/constants/mapstate"
	"kaisu/pokemon/constants/playerstate"
	"kaisu/pokemon/src/building"
	"kaisu/pokemon/src/console"
	"kaisu/pokemon/src/mapparser"
	"kaisu/pokemon/src/player"
	. "kaisu/pokemon/src/utils"

	tl "github.com/JoelOtter/termloop"
)

func loadMapLevel(fileurl string) {

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

	// render player
	playerstate.PLAYER = &player.Player{
		Entity: tl.NewEntity(1, 1, 1, 1),
	}

	// Set the character at position (0, 0) on the entity.
	playerstate.PLAYER.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '@'})
	constants.CURLEVEL.AddEntity(playerstate.PLAYER)
}

func main() {

	constants.GAME = tl.NewGame()
	constants.GAME.Screen().SetFps(30)
	mapstate.LOCATIONS = make(map[string]*building.Building)

	// level.AddEntity(tl.NewRectangle(10, 10, 10, 20, tl.ColorBlue))
	// level.AddEntity(tl.NewText(10, 10, "House", tl.ColorRed, tl.ColorWhite))
	// l := tl.NewBaseLevel(tl.Cell{Bg: 76, Fg: 1})
	// fmt.Println(len(level.Entities))

	loadMapLevel("assets/maps/town1.json")

	constants.CONSOLE_TEXT = console.NewConsoleText(constants.GAME.Screen(), constants.CURLEVEL)
	constants.CONSOLE_TEXT.SetText(constants.DISPLAYED_TEXT)

	constants.GAME.Screen().SetLevel(constants.CURLEVEL)
	constants.GAME.Start()
}
