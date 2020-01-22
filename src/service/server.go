package service

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"../finder"
)

// Logly is a type that implements the Retrieve() "method"
type Logly struct {
	logFile string
}

// GetLogs returns all the logs for this server.
func (l *Logly) GetLogs(yes int, logs *string) error {
	b, err := ioutil.ReadFile((*l).logFile)
	if err != nil {
		return err
	}

	*logs = string(b)
	return nil
}

// GrepLogs greps logs and returns the corresponding results.
func (l *Logly) GrepLogs(expression string, results *string) error {
	*results = finder.Find((*l).logFile, expression)

	return nil
}

// Server is a long running function that accepts new RPC requests and fufills them
func Server(logFile string) {
	var logly *Logly = &Logly{logFile}
	err := rpc.Register(logly)

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
	log.Printf("Serving RPC server on port %d for log file: %s", 8080, logFile)

	// Start accept incoming HTTP connections
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving:", err)
	}
}
