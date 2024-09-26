package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B float64
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1121")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &Args{A: 3.5, B: 2.2}
	var result float64

	err = client.Call("MultiplyService.Multiply", args, &result)
	if err != nil {
		log.Fatal("error calling Multiply:", err)
	}

	fmt.Printf("Result: %.2f\n", result)
}
