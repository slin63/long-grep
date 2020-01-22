package service

import (
	"fmt"
	"log"
	"net/rpc"
	"sync"

	"../config"
)

// Client for querying.
func Client(expressionPtr *string) {
	var wg sync.WaitGroup

	addresses, err := config.Addresses()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(addresses)
	for _, address := range addresses {
		log.Println(address)
		wg.Add(1)
		go grepLogs(address, expressionPtr, &wg)
	}

	wg.Wait()
}

func grepLogs(address string, expressionPtr *string, wg *sync.WaitGroup) string {
	defer wg.Done()
	log.Println("Dispatching RPC for", address)
	var logs string
	client, err := rpc.DialHTTP("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	client.Call("Logly.GrepLogs", *expressionPtr, &logs)
	return logs
}
