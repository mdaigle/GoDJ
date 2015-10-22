package control

import (
	//"spotify"
	format "fmt"
	"github.com/mdaigle/GoDJ/perform"
)

func ProcessCommand(command string) bool {
	switch command {
	case "p" :
		perform.PlayPause()
	case "n" :
		perform.Next()
	case "b" :
		perform.Previous()
	case "i" :
		perform.IncreaseVolume()
	case "d" :
		perform.DecreaseVolume()
//	case "b", "browse" :
//	//display music menu
//	case "c", "custom" :
//	//display prompts to specify custom mood
	case "q", "quit" :
		return true;
	default:
		format.Print("Invalid Command")
	}
	return false;

}