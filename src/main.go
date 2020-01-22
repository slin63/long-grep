package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"

	"./service"
	"./setup"
)

func main() {
	isClientPtr := flag.Bool("client", false, "configure process as client")
	isServerPtr := flag.Bool("server", false, "configure process as server")
	expressionPtr := flag.String("expression", "", "regular expression")
	lognamePtr := flag.String(
		"logname",
		fmt.Sprintf("machine.%s.log", os.Getenv("MACHINEID")),
		"name of the generated log file",
	)

	flag.Parse()

	switch {
	case *isServerPtr:
		setup.Setup(*lognamePtr)
		service.Server(*lognamePtr)

	case *isClientPtr:
		if *expressionPtr == "" {
			log.Fatalln("Must specify expression if using as a client")
		}
		// Validate the regular expression
		_, err := regexp.Compile(*expressionPtr)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(*expressionPtr)

		service.Client(expressionPtr)
	default:
		log.Fatalln("Usage: main [-client] [-server] [--expression=<regular expression>]")
	}
}
