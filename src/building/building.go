package building

import tl "github.com/JoelOtter/termloop"

type Building struct {
	// *tl.Entity
	Shape *tl.Rectangle
	title *tl.Text
	level *tl.BaseLevel
}

func InitBuilding(level *tl.BaseLevel) *Building {
	aBuilding := Building{
		// Entity: tl.NewEntity(1, 1, 1, 1),
		Shape: tl.NewRectangle(10, 10, 10, 5, tl.ColorBlue),
		title: tl.NewText(10, 10, "House", tl.ColorRed, tl.ColorWhite),
		level: level,
	}

	// level.AddEntity(pokemonCenter.Entity)
	level.AddEntity(aBuilding.Shape)
	level.AddEntity(aBuilding.title)

	return &aBuilding
}
