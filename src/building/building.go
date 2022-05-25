package building

import tl "github.com/JoelOtter/termloop"

type Building struct {
	*tl.Entity
	title *tl.Text
	shape *tl.Rectangle
	level *tl.BaseLevel
}

func InitBuilding(level *tl.BaseLevel) {
	pokemonCenter := Building{
		// Entity: tl.NewEntity(1, 1, 1, 1),
		title: tl.NewText(10, 10, "House", tl.ColorRed, tl.ColorWhite),
		shape: tl.NewRectangle(10, 10, 10, 5, tl.ColorBlue),
		level: level,
	}

	// level.AddEntity(pokemonCenter.Entity)
	level.AddEntity(pokemonCenter.shape)
	level.AddEntity(pokemonCenter.title)
}
