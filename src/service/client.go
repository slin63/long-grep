package service

import (
	"fmt"
	"log"
	"net/rpc"
	"sync"

	"../config"
)

// Client for querying.
func Client() {
	var wg sync.WaitGroup

	addresses, err := config.Addresses()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(addresses)
	for _, address := range addresses {
		log.Println(address)
		wg.Add(1)
		go getLogs(address, &wg)
	}

	wg.Wait()
}

func getLogs(address string, wg *sync.WaitGroup) string {
	defer wg.Done()
	log.Println("Dispatching RPC for", address)
	var logs string
	client, err := rpc.DialHTTP("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	client.Call("Logly.GetLogs", 1, &logs)
	fmt.Println(logs)
	return logs
}
