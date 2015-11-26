package view
import "fmt"

func DisplayMainMenu() {
	fmt.Println("\nb - previous track");
	fmt.Println("d - decrease volume");
	fmt.Println("i - increase volume");
	fmt.Println("n - next track");
	fmt.Println("p - play/pause");
	fmt.Println("b, browse - browse profiles");
//	display music menu
//	case "c", "custom" :
//	//display prompts to specify custom mood
	fmt.Println("q, quit - quit\n");
}