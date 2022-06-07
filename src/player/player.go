package player

import (
	"fmt"
	"kaisu/pokemon/constants"
	"kaisu/pokemon/constants/mapstate"
	"kaisu/pokemon/src/building"
	"kaisu/pokemon/src/maploader"

	. "kaisu/pokemon/src/utils"
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

func (player *Player) SetPrevPosition() {
	player.SetPosition(player.prevX, player.prevY)
}

func (player *Player) Collide(collision tl.Physical) {

	collisionObject := fmt.Sprintf("%+v\n", collision)

	switch collision.(type) {
	case building.Warppoint:
		// constants.CONSOLE_TEXT.SetText("collided with warpPoint: " + collisionObject)
		collidedWarppoint, _ := collision.(building.Warppoint)
		Use(collidedWarppoint)
		constants.CONSOLE_TEXT.SetText("collided with warpPoint going to url: " + collidedWarppoint.Url)
		maploader.LoadRoom(collidedWarppoint.Url)
		player.SetPosition(1, 1)
		constants.CURLEVEL.AddEntity(player)
		constants.GAME.Screen().SetLevel(constants.CURLEVEL)

	case building.Building:
		constants.CONSOLE_TEXT.SetText("collided with building: " + collisionObject)

		// block from moving
		player.SetPrevPosition()
		constants.CONSOLE_TEXT.SetPrevPosition()

		// locate the building structure, only if collided
		for matchingName, matchingBuilding := range mapstate.LOCATIONS {

			collidedBuilding, _ := collision.(building.Building)
			if collidedBuilding.Name == matchingBuilding.Name {
				constants.CONSOLE_TEXT.SetText("collided with building name: " + matchingName)
			}

			// in the future, make an extra check so it only searches for buildings in the current map/level
		}
	default:
		constants.CONSOLE_TEXT.SetText("collided with something: " + collisionObject)
	}

	// if _, ok := collision.(*tl.Rectangle); ok {
	//   player.SetPosition(player.prevX, player.prevY)
	//   constants.CONSOLE_TEXT.SetText("collided with rectangle")
	//   // locate the building structure, only if collided
	//   for key, element := range mapstate.LOCATIONS {
	//     foundElement := (collision.(*tl.Rectangle) == element.Rectangle)
	//     if foundElement {
	//       fmt.Println(key, element)
	//     }
	//   }
	// }
}
