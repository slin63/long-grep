package main

import (
	"flag"
	"fmt"
	"log"
	"os"

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
	// configFilePtr := flag.String("config", "", "Location of Config File")

	flag.Parse()

	// os.Setenv("CONFIG", *configFilePtr)

	switch {
	case *isServerPtr:
		setup.Setup(*lognamePtr)
		service.Server(*lognamePtr)

	case *isClientPtr:
		// 	if *expressionPtr == "" {
		// 		log.Fatalln("Must specify expression if using as a client")
		// 	}
		// 	_, err := regexp.Compile(*expressionPtr)
		// 	if err != nil {
		// 		log.Fatalln(err)
		// 	}
		// 	service.Client(*expressionPtr)
		log.Println(*expressionPtr)
		service.Client()
	default:
		log.Fatalln("Usage: main [-client] [-server] [--expression=<regular expression>]")
	}
}
