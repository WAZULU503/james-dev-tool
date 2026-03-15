package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/WAZULU503/Name-james-dev-tool/internal/ports"
	"github.com/WAZULU503/Name-james-dev-tool/internal/ui"
)

const version = "v0.1.0"

func main() {

	if len(os.Args) < 2 {
		printUsage()
		return
	}

	if os.Args[1] == "version" {
		fmt.Println("jdt", version)
		return
	}

	if len(os.Args) < 3 {
		printUsage()
		return
	}

	cmd := os.Args[1]
	sub := os.Args[2]

	switch cmd {

	case "p", "port":

		switch sub {

		case "ls", "list":

			items, err := ports.GetPorts()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			ui.PrintTable(items)

		case "k", "kill":

			if len(os.Args) < 4 {
				fmt.Println("Specify port")
				return
			}

			port, _ := strconv.Atoi(os.Args[3])

			err := ports.KillPort(port)

			if err != nil {
				fmt.Println(err)
			}

		case "f", "free":

			err := ports.FreePorts()

			if err != nil {
				fmt.Println(err)
			}

		default:
			fmt.Println("Unknown port command")
		}

	default:
		printUsage()
	}
}

func printUsage() {

	fmt.Println("jdt (James Dev Tool)")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  jdt p ls        List active dev ports")
	fmt.Println("  jdt p k <port>  Kill process on port")
	fmt.Println("  jdt p f         Free all dev ports")
	fmt.Println("  jdt version     Show version")
}
