package main

import (
	"LaraGo/lara"
	"errors"
	"github.com/fatih/color"
	"log"
	"os"
)

var la lara.Lara

const version = "1.0.0"

func main() {
	arg1, arg2, arg3, err := validateInput()
	if err != nil {
		exitGracefully(err)
	}

	switch arg1 {
	case "help":
		showHelp()
	case "version":
		color.Yellow(version)
	default:
		log.Println(arg2, arg3)
	}
}

func exitGracefully(err error, msg ...string) {
	message := ""
	if len(msg) > 0 {
		message = msg[0]
	}

	if err != nil {
		color.Red("%s\n%s", message, err)
	}

	if len(message) > 0 {
		color.Yellow(message)
	} else {
		color.Green("Finished")
	}

	os.Exit(0)
}

func validateInput() (string, string, string, error) {
	var arg1, arg2, arg3 string
	if len(os.Args) > 1 {
		arg1 = os.Args[1]

		if len(os.Args) >= 3 {
			arg2 = os.Args[2]
		}

		if len(os.Args) >= 4 {
			arg3 = os.Args[3]
		}
	} else {
		color.Red("Error: at least one argument is required")
		showHelp()
		return "", "", "", errors.New("at least one argument is required")
	}

	return arg1, arg2, arg3, nil
}

func showHelp() {
	color.Yellow(`Usage:
    help  - print this help message
    version - print version number
`)
}

func exitGraceFully(err error, msg ...string) {
	message := ""
	if len(msg) > 0 {
		message = msg[0]
	}

	if err != nil {
		color.Red("%s\n%s", message, err)
	}

	if len(message) > 0 {
		color.Yellow(message)
	} else {
		color.Green("Finished")
	}

	os.Exit(0)
}
