package building

import (
	"kaisu/pokemon/constants"

	tl "github.com/JoelOtter/termloop"
)

type Building struct {
	// *tl.Entity
	// Shape *tl.Rectangle
	*tl.Rectangle
	Title *tl.Text
	Name  string
}

func NewBuilding(data map[string]interface{}) *Building {
	xPos := int(data["x"].(float64))
	yPos := int(data["y"].(float64))
	width := int(data["width"].(float64))
	height := int(data["height"].(float64))
	color := tl.Attr(data["color"].(float64))
	// name := string(data["name"].(float64))
	name := string(data["name"].(string))

	// Entity: tl.NewEntity(1, 1, 1, 1),
	// Shape: tl.NewRectangle(xPos, yPos, width, height, tl.ColorBlue),
	aBuilding := Building{
		// Shape: tl.NewRectangle(xPos, yPos, width, height, color),
		Rectangle: tl.NewRectangle(xPos, yPos, width, height, color),
	}

	aBuilding.Title = tl.NewText(xPos, yPos, name, tl.ColorRed, tl.ColorWhite)
	aBuilding.Name = name

	// level.AddEntity(pokemonCenter.Entity)
	constants.CURLEVEL.AddEntity(aBuilding)
	constants.CURLEVEL.AddEntity(aBuilding.Title)

	return &aBuilding
}

// func InitBuilding(level *tl.BaseLevel) *Building {
// }
