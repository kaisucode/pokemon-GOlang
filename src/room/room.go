package room

import (
	"kaisu/pokemon/constants"

	tl "github.com/JoelOtter/termloop"
)

type Room struct {
	*tl.Rectangle
	Title *tl.Text
	Name  string
}

func NewRoom(data map[string]interface{}) *Room {
	xPos := int(data["x"].(float64))
	yPos := int(data["y"].(float64))
	width := int(data["width"].(float64))
	height := int(data["height"].(float64))
	color := tl.Attr(data["color"].(float64))
	name := string(data["name"].(string))

	// Entity: tl.NewEntity(1, 1, 1, 1),
	// Shape: tl.NewRectangle(xPos, yPos, width, height, tl.ColorBlue),
	aRoom := Room{
		Rectangle: tl.NewRectangle(xPos, yPos, width, height, color),
	}

	titleX := width/2 + xPos - len(name)/2
	titleY := height/2 + yPos
	aRoom.Title = tl.NewText(titleX, titleY, name, tl.ColorRed, tl.ColorWhite)
	aRoom.Name = name

	constants.CURLEVEL.AddEntity(aRoom)
	constants.CURLEVEL.AddEntity(aRoom.Title)

	return &aRoom
}
