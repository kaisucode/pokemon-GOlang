package player

import (
	"fmt"
	"kaisu/pokemon/constants"
	"kaisu/pokemon/constants/mapstate"
	"kaisu/pokemon/src/building"
	"strconv"

	tl "github.com/JoelOtter/termloop"
)

type Player struct {
	*tl.Entity
	prevX, prevY int
}

func (player *Player) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()
	x, y := player.Position()
	constants.CURLEVEL.SetOffset(screenWidth/2-x, screenHeight/2-y)
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

	// temp, ok := collision.(*building.Building)
	// huh := collision.(*building.Building)
	// constants.CONSOLE_TEXT.SetText("huh: " + string(huh))
	// constants.CONSOLE_TEXT.SetText("temp: " + string(temp))
	// constants.CONSOLE_TEXT.SetText("ok: %t", ok)

	collisionObject := fmt.Sprintf("%+v\n", collision)

	switch collision.(type) {
	case building.Building:
		constants.CONSOLE_TEXT.SetText("collided with building: " + collisionObject)
		player.SetPosition(player.prevX, player.prevY)
		// block from moving

		// locate the building structure, only if collided
		for buildingName, element := range mapstate.LOCATIONS {

			// temp, ok := collision.(*building.Building)
			// use(temp, ok)
			// foundElement := (collision.(*building.Building) == element)
			foundElement := (buildingName == element.Name)

			// in the future, make an extra check so it only searches for buildings in the current map/level
			if foundElement {
				constants.CONSOLE_TEXT.SetText("collided with building name: " + buildingName)
			}
		}
	default:
		constants.CONSOLE_TEXT.SetText("collided with something: " + collisionObject)
	}

	if _, ok := collision.(*tl.Rectangle); ok {
		player.SetPosition(player.prevX, player.prevY)
		constants.CONSOLE_TEXT.SetText("collided with rectangle")
		// locate the building structure, only if collided
		for key, element := range mapstate.LOCATIONS {
			foundElement := (collision.(*tl.Rectangle) == element.Rectangle)
			if foundElement {
				fmt.Println(key, element)
			}
		}
	}
}
