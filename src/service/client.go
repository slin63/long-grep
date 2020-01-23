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

	go func() {
		wg.Wait()
		close(logs)
	}()

	counter := 0
	for msg := range logs {
		fmt.Println(msg)
		counter++
	}

	log.Printf("Received %d messages.", counter)
}

func grepLogs(address string, expressionPtr *string, wg *sync.WaitGroup, logs chan string) {
	defer wg.Done()
	log.Println("Dispatching RPC for", address)
	var reply string
	client, err := rpc.DialHTTP("tcp", address)
	if err != nil {
		log.Println("Error - GrepLogs:", err)
	} else {
		client.Call("Logly.GrepLogs", *expressionPtr, &reply)
		log.Println("done:", address, len(reply))
		logs <- reply
	}

}
