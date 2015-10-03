package main

import(
	"bufio"
	"github.com/mdaigle/GoDJ/control"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Welcome to GoDJ!\n")
	for {
		fmt.Print("Enter command: ")
		command, err := r.ReadString('\n')
		if err != nil {
			fmt.Print("Whoops, didn't quite catch that, please modify your" +
			"command.")
		}
		command = strings.TrimSpace(command)
		end := control.ProcessCommand(command)
		if end {
			break
		}
	}
	fmt.Print("Come back soon!")
}