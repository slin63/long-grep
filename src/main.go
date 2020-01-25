package main

import (
	"flag"
	"log"
	"regexp"

	"./service"
	"./setup"
)

func main() {
	isClientPtr := flag.Bool("client", false, "configure process as client")
	isServerPtr := flag.Bool("server", false, "configure process as server")
	generateLogs := flag.Bool("generate", false, "generate logs at each server")
	expressionPtr := flag.String("expression", "", "regular expression")
	lognamePtr := flag.String("logname", "", "name of the log file to grep")

	flag.Parse()

	switch {
	case *isServerPtr:
		if *generateLogs {
			setup.Setup(*lognamePtr)
		}
		service.Server(*lognamePtr)

	case *isClientPtr:
		if *expressionPtr == "" {
			log.Fatalln("Must specify expression if using as a client")
		}
		if *lognamePtr == "" {
			log.Fatalln("Must specify logname if using as a client")
		}
		// Validate the regular expression
		_, err := regexp.Compile(*expressionPtr)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(*expressionPtr)

		service.Client(expressionPtr)
	default:
		log.Fatalln("Usage: main [-client] [-server] [-generate] [--logname=<name of log file>] [--expression=<regular expression>]")
	}
}
