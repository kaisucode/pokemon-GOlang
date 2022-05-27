package main

import (
	"io/ioutil"
	"kaisu/pokemon/constants"
	"kaisu/pokemon/constants/playerstate"
	"kaisu/pokemon/src/building"
	"kaisu/pokemon/src/console"
	"kaisu/pokemon/src/player"

	tl "github.com/JoelOtter/termloop"
)

type Player struct {
	*tl.Entity
}

// Here we define a parse function for reading a Player out of JSON.
func parsePlayer(data map[string]interface{}) tl.Drawable {
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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	constants.GAME = tl.NewGame()
	constants.GAME.Screen().SetFps(30)
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
		Ch: '.',
	})

	// level.AddEntity(tl.NewRectangle(10, 10, 10, 20, tl.ColorBlue))

	// level.AddEntity(tl.NewText(10, 10, "House", tl.ColorRed, tl.ColorWhite))

	// l := tl.NewBaseLevel(tl.Cell{Bg: 76, Fg: 1})
	lmap, err := ioutil.ReadFile("levels/town1.json")
	checkErr(err)
	parsers := make(map[string]tl.EntityParser)
	parsers["Player"] = parsePlayer
	err = tl.LoadLevelFromMap(string(lmap), parsers, level)
	checkErr(err)

	constants.CONSOLE_TEXT = console.NewConsoleText(constants.GAME.Screen(), level)
	constants.CONSOLE_TEXT.SetText(constants.DISPLAYED_TEXT)
	// consoleText.SetText("new text!!")
	// fmt.Println(len(level.Entities))

	locations := make(map[string]*building.Building)
	pokemonCenter := building.InitBuilding(level)
	locations["pokemonCenter"] = pokemonCenter

	playerstate.PLAYER = &player.Player{
		Entity:    tl.NewEntity(1, 1, 1, 1),
		Locations: &locations,
		Level:     level,
	}

	// Set the character at position (0, 0) on the entity.
	playerstate.PLAYER.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '@'})
	level.AddEntity(playerstate.PLAYER)
	constants.GAME.Screen().SetLevel(level)

	constants.GAME.Start()
}
