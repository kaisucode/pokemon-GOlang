package main

import (
	"io/ioutil"
	"kaisu/pokemon/constants"
	"kaisu/pokemon/constants/playerstate"
	"kaisu/pokemon/src/building"
	"kaisu/pokemon/src/console"
	"kaisu/pokemon/src/mapparser"
	"kaisu/pokemon/src/player"
	. "kaisu/pokemon/src/utils"

	tl "github.com/JoelOtter/termloop"
)

func loadMap(fileurl string) *tl.BaseLevel {

	// create a base level
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
		Ch: '.',
	})

	lmap, err := ioutil.ReadFile(fileurl)
	CheckErr(err)

	parsers := make(map[string]tl.EntityParser)
	parsers["Player"] = mapparser.ParsePlayer
	parsers["Buildings"] = mapparser.ParseBuildings // handles collision rects and warp points
	// parsers["Npc"] = mapparser.ParsePlayer
	// parsers["Grass"] = mapparser.ParsePlayer
	err = tl.LoadLevelFromMap(string(lmap), parsers, level)
	CheckErr(err)

	// render buildings (should be done in building init)
	pokemonCenter := building.InitBuilding(level)
	constants.LOCATIONS["pokemonCenter"] = pokemonCenter

	// render player
	playerstate.PLAYER = &player.Player{
		Entity: tl.NewEntity(1, 1, 1, 1),
		Level:  level,
	}

	// Set the character at position (0, 0) on the entity.
	playerstate.PLAYER.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '@'})
	level.AddEntity(playerstate.PLAYER)

	return level
}

func main() {

	constants.GAME = tl.NewGame()
	constants.GAME.Screen().SetFps(30)
	constants.LOCATIONS = make(map[string]*building.Building)

	// level.AddEntity(tl.NewRectangle(10, 10, 10, 20, tl.ColorBlue))
	// level.AddEntity(tl.NewText(10, 10, "House", tl.ColorRed, tl.ColorWhite))
	// l := tl.NewBaseLevel(tl.Cell{Bg: 76, Fg: 1})
	// consoleText.SetText("new text!!")
	// fmt.Println(len(level.Entities))

	level := loadMap("assets/maps/town1.json")

	constants.CONSOLE_TEXT = console.NewConsoleText(constants.GAME.Screen(), level)
	constants.CONSOLE_TEXT.SetText(constants.DISPLAYED_TEXT)

	constants.GAME.Screen().SetLevel(level)
	constants.GAME.Start()
}
