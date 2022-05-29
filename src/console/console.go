package console

import (
	"kaisu/pokemon/constants/logger"

	tl "github.com/JoelOtter/termloop"
)

type ConsoleText struct {
	*tl.Text
	logs         []string
	prevX, prevY int
}

func (consoleText *ConsoleText) AddMessage(newMessage string) {
	// consoleText.logs = append(consoleText.logs, newMessage)
	logger.INFO.Println(newMessage)
}

func (consoleText *ConsoleText) SetText(newMessage string) {
	// consoleText.logs = append(consoleText.logs, newMessage)
	consoleText.Text.SetText(newMessage)
	logger.INFO.Println(newMessage)
}

func (consoleText *ConsoleText) Draw(screen *tl.Screen) {
	consoleText.Text.Draw(screen)
}

func (consoleText *ConsoleText) SetPrevPosition() {
	consoleText.SetPosition(consoleText.prevX, consoleText.prevY)
}

func (consoleText *ConsoleText) Tick(ev tl.Event) {
	// Enable arrow key movement
	if ev.Type == tl.EventKey {
		// x, y := m.Position()
		consoleText.prevX, consoleText.prevY = consoleText.Position()

		switch ev.Key {
		case tl.KeyArrowRight:
			consoleText.SetPosition(consoleText.prevX+1, consoleText.prevY)
		case tl.KeyArrowLeft:
			consoleText.SetPosition(consoleText.prevX-1, consoleText.prevY)
		case tl.KeyArrowUp:
			consoleText.SetPosition(consoleText.prevX, consoleText.prevY-1)
		case tl.KeyArrowDown:
			consoleText.SetPosition(consoleText.prevX, consoleText.prevY+1)
		}

		curX, curY := consoleText.Position()

		consoleText.SetPosition(curX, curY)
	}
}

func NewConsoleText(screen *tl.Screen, level *tl.BaseLevel) *ConsoleText {

	_, screenHeight := screen.Size()
	// x, y := player.Position()
	// player.Level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	// game.Screen().AddEntity(&consoleText)

	// x, y := player.Position()
	// player.Level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	// level.AddEntity(&consoleText)

	// xPos := screenWidth / 2
	// yPos := screenHeight / 2
	xPos := -115
	yPos := screenHeight/2 + 24

	textToDisplay := "HELLO WORLD"
	consoleText := ConsoleText{
		tl.NewText(xPos, yPos, textToDisplay, tl.ColorWhite, tl.ColorBlack),
		make([]string, 10),
		0, 0,
	}
	level.AddEntity(&consoleText)

	return &consoleText
}
