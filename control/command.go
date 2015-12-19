package control

import (
//	"github.com/rapito/go-spotify/spotify"
	format "fmt"
	"github.com/mdaigle/GoDJ/perform"
	"github.com/mdaigle/GoDJ/view"
)

func ProcessCommand(command string) bool {
	switch command {
	case "b" :
		perform.Previous()
	case "d" :
		perform.DecreaseVolume()
	case "i" :
		perform.IncreaseVolume()
	case "h", "help" :
		view.DisplayMainMenu()
	case "n" :
		perform.Next()
	case "p" :
		perform.PlayPause()
//	case "b", "browse" :
//	display music menu
//	case "c", "custom" :
//	//display prompts to specify custom mood
	case "q", "quit" :
		return true;
	default:
		format.Print("Invalid Command")
	}
	return false;

}