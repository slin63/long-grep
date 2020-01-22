package service

import (
	"fmt"
	"log"
	"net/rpc"
	"sync"

	"../config"
)

// TODO
// - Add timeout for dead machines

// Client for querying logs
func Client(expressionPtr *string) {
	logs := make(chan string)
	results := make(chan string)
	var wg sync.WaitGroup

	addresses, err := config.Addresses()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(addresses)
	for _, address := range addresses {
		log.Println(address)
		wg.Add(1)
		go grepLogs(address, expressionPtr, &wg, logs)
	}
	go compileLogs(logs, results)

	wg.Wait()
	close(logs)
	fmt.Println(<-results)
}

func compileLogs(logs, results chan string) {
	var compiled, log string
	ok := true
	for ok {
		log, ok = <-logs
		compiled += log
	}
	results <- compiled
}

func grepLogs(address string, expressionPtr *string, wg *sync.WaitGroup, logs chan string) {
	defer wg.Done()
	log.Println("Dispatching RPC for", address)
	var reply string
	client, err := rpc.DialHTTP("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	client.Call("Logly.GrepLogs", *expressionPtr, &reply)
	log.Println("done:", address, len(reply))
	logs <- reply
}
