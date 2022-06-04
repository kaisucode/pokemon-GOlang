package building

import (
	"kaisu/pokemon/constants"

	tl "github.com/JoelOtter/termloop"
)

type Warppoint struct {
	// *tl.Entity
	*tl.Rectangle
	url string
}

func NewWarppoint(xPos int, yPos int, url string) *Warppoint {
	// xPos := int(data["x"].(float64))
	// yPos := int(data["y"].(float64))
	// width := int(data["width"].(float64))
	// height := int(data["height"].(float64))
	// color := tl.Attr(data["color"].(float64))
	// name := string(data["name"].(string))

	aWarppoint := Warppoint{
		Rectangle: tl.NewRectangle(xPos, yPos, 1, 1, tl.ColorBlue),
		url:       url,
	}
	// Entity: tl.NewEntity(1, 1, 1, 1),
	// Shape: tl.NewRectangle(xPos, yPos, width, height, tl.ColorBlue),
	// aBuilding := Building{
	//   Rectangle: tl.NewRectangle(xPos, yPos, width, height, color),
	// }

	// titleX := width/2 + xPos - len(name)/2
	// titleY := height/2 + yPos
	// aBuilding.Title = tl.NewText(titleX, titleY, name, tl.ColorRed, tl.ColorWhite)
	// aBuilding.Name = name

	constants.CURLEVEL.AddEntity(aWarppoint)
	// constants.CURLEVEL.AddEntity(aBuilding.Title)

	return &aWarppoint
}
