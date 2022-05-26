package console

import (
	tl "github.com/JoelOtter/termloop"
)

type ConsoleText struct {
	*tl.Text
	Level *tl.BaseLevel
}

func (consoleText *ConsoleText) Draw(screen *tl.Screen) {
	// screenWidth, screenHeight := screen.Size()
	// x, y := consoleText.Position()
	// consoleText.Level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	consoleText.Text.Draw(screen)
}

// func (m *ConsoleText) Tick(ev tl.Event) {
//   // Enable arrow key movement
//   if ev.Type == tl.EventKey {
//     x, y := m.Position()
//     switch ev.Key {
//     case tl.KeyArrowRight:
//       x += 1
//     case tl.KeyArrowLeft:
//       x -= 1
//     case tl.KeyArrowUp:
//       y -= 1
//     case tl.KeyArrowDown:
//       y += 1
//     }
//     m.SetPosition(x, y)
//   }
// }
