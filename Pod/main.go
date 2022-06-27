package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		panic(err)
	}

	defer lis.Close()

	for {
		_, err := lis.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		fmt.Println("hello golang application")
	}
}
