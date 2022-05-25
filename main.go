package main

import (
	"kaisu/pokemon/src/building"
	"kaisu/pokemon/src/player"

	tl "github.com/JoelOtter/termloop"
)

func main() {
	game := tl.NewGame()
	game.Screen().SetFps(30)
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
		Ch: '.',
	})

	// level.AddEntity(tl.NewRectangle(10, 10, 10, 20, tl.ColorBlue))

	// level.AddEntity(tl.NewText(10, 10, "House", tl.ColorRed, tl.ColorWhite))
	building.InitBuilding(level)

	player := player.Player{
		Entity: tl.NewEntity(1, 1, 1, 1),
		Level:  level,
	}
	// Set the character at position (0, 0) on the entity.
	player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '옷'})
	level.AddEntity(&player)
	game.Screen().SetLevel(level)
	game.Start()
}
