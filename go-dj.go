package main

import(
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command: ")
		command, err := r.ReadString('\n')
		command = strings.TrimSpace(command)
		if err != nil {
			fmt.Println("Invalid command.")
		}
		if command == "end" {
			break
		}
	}
}