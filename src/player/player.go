package player

import (
	"fmt"
	"kaisu/pokemon/constants"
	"strconv"

	tl "github.com/JoelOtter/termloop"
)

type Player struct {
	*tl.Entity
	Level        *tl.BaseLevel
	prevX, prevY int
}

func (player *Player) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()
	x, y := player.Position()
	player.Level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	player.Entity.Draw(screen)
}

func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		player.prevX, player.prevY = player.Position()
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			player.SetPosition(player.prevX+1, player.prevY)
		case tl.KeyArrowLeft:
			player.SetPosition(player.prevX-1, player.prevY)
		case tl.KeyArrowUp:
			player.SetPosition(player.prevX, player.prevY-1)
		case tl.KeyArrowDown:
			player.SetPosition(player.prevX, player.prevY+1)
		}

		curX, curY := player.Position()
		constants.CONSOLE_TEXT.SetText("(" + strconv.Itoa(curX) + ", " + strconv.Itoa(curY) + ")")
	}
}

func (player *Player) Collide(collision tl.Physical) {
	// Check if it's a Rectangle we're colliding with
	// fmt.Println(collision)
	// fmt.Println("--------")

	// block from moving
	if _, ok := collision.(*tl.Rectangle); ok {
		player.SetPosition(player.prevX, player.prevY)

		// locate the building structure, only if collided
		for key, element := range constants.LOCATIONS {

			foundElement := (collision.(*tl.Rectangle) == element.Shape)

			// in the future, make an extra check so it only searches for buildings in the current map/level
			if foundElement {
				fmt.Println(key, element)
			}
		}
	}
}
