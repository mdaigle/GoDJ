package control

import (
//"spotify"
	"fmt"
)

func ProcessCommand(command string) bool {
	switch command {
	case "p", "play" :
	//start music streaming
	case "s", "pause" :
	//stop music streaming
	case "b", "browse" :
	//display music menu
	case "c", "custom" :
	//display prompts to specify custom mood
	case "q", "quit" :
		return true;
	default:
		fmt.Print("Invalid Command")
	}
	return false;

}