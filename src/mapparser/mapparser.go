package mapparser

import (
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
