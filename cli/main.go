package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tyomhk2015/gocoin/goserver/explorer"
	"github.com/tyomhk2015/gocoin/restapi/rest"
)

func main() {
	if !checkLength() {
		// Show introduction when there are no command line arguments.
		introduce()
	}

	// Excecute commands
	portFlag := flag.Int("port", 8888, "Set the port number of the server.")
	mode := flag.String("mode", "explorer", "Choose the way to view the blockchain, 'html' or 'rest'")
	flag.Parse()

	fmt.Println(*mode, *portFlag)

	switch *mode {
	case "html":
		explorer.Start(*portFlag)
	case "rest":
		rest.Start(*portFlag)
	case "both":
		go explorer.Start(*portFlag)
		rest.Start(*portFlag + 1)
	default:
		introduce()
	}
}

func introduce() {
	welcome := "Welcome to gocoin CLI program. 💻\nSelect desired program to initiate."
	programs := "\n\n👉 -mode=rest: Choose between 'html', 'rest', and 'both'.\n👉 -port=8888: Set port number of the server.\n"
	fmt.Println(welcome, programs)
	os.Exit(0)
}

func checkLength() bool {
	// Check the length of the command-line arguements.
	// If the length is longer than 1, return true.
	if len(os.Args) < 2 {
		return false
	}
	return true
}

func setPort() {
	portFlag := flag.Int("port", 8888, "Set the port number of the server.")
	fmt.Println(*portFlag)
}

func setMode() {
	mode := flag.String("mode", "explorer", "Choose the way to view the blockchain, 'html' or 'rest'")
	fmt.Println(*mode)
}
