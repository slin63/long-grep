package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

type Task int

var logFile = os.Args[1]

func (t *Task) getLogs(_, logs *string) error {
	b, err := ioutil.ReadFile(logFile)
	if err != nil {
		return err
	}

	*logs = string(b)
	return nil
}

func main() {
	task := new(Task)
	err := rpc.Register(task)
	if err != nil {
		panic(err)
	}

	// Register an HTTP handler
	rpc.HandleHTTP()

	// Listen to TCP connections on port 8080
	listener, e := net.Listen("tcp", ":8080")
	if e != nil {
		log.Fatal("Listen error:", e)
	}
	log.Printf("Serving RPC server on port %d", 8080)

	// Start accept incoming HTTP connections
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving:", err)
	}
}
