package console

import (
	tl "github.com/JoelOtter/termloop"
)

type ConsoleText struct {
	*tl.Text
}

func (consoleText *ConsoleText) Draw(screen *tl.Screen) {
	consoleText.Text.Draw(screen)
}

func (m *ConsoleText) Tick(ev tl.Event) {
	// Enable arrow key movement
	if ev.Type == tl.EventKey {
		x, y := m.Position()
		switch ev.Key {
		case tl.KeyArrowRight:
			x += 1
		case tl.KeyArrowLeft:
			x -= 1
		case tl.KeyArrowUp:
			y -= 1
		case tl.KeyArrowDown:
			y += 1
		}
		m.SetPosition(x, y)
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
	consoleText := ConsoleText{tl.NewText(xPos, yPos, textToDisplay, tl.ColorWhite, tl.ColorBlack)}
	level.AddEntity(&consoleText)

	return &consoleText
}
