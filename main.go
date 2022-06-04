package main

import (
	"kaisu/pokemon/constants"
	"kaisu/pokemon/constants/logger"
	"kaisu/pokemon/constants/mapstate"
	"kaisu/pokemon/src/building"
	"kaisu/pokemon/src/console"
	"kaisu/pokemon/src/maploader"
	"log"
	"os"

	// . "kaisu/pokemon/src/utils"

	tl "github.com/JoelOtter/termloop"
)

func main() {

	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	// log.SetOutput(f)

	logger.INFO = log.New(f, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	logger.INFO.Println("hello")

	constants.GAME = tl.NewGame()
	constants.GAME.Screen().SetFps(30)
	mapstate.LOCATIONS = make(map[string]*building.Building)

	// level.AddEntity(tl.NewRectangle(10, 10, 10, 20, tl.ColorBlue))
	// level.AddEntity(tl.NewText(10, 10, "House", tl.ColorRed, tl.ColorWhite))
	// l := tl.NewBaseLevel(tl.Cell{Bg: 76, Fg: 1})
	// fmt.Println(len(level.Entities))

	maploader.LoadMapLevel("assets/maps/town1.json")
	// maploader.LoadRoom("assets/maps/pkmncenter1.json")

	constants.CONSOLE_TEXT = console.NewConsoleText(constants.GAME.Screen(), constants.CURLEVEL)
	constants.CONSOLE_TEXT.SetText(constants.DISPLAYED_TEXT)

	constants.GAME.Screen().SetLevel(constants.CURLEVEL)
	constants.GAME.Start()
}
