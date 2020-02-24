package main

import (
	"flag"
	"fmt"
	"net"
	"sort"
)

var targetip string

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", targetip, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	var targetPortCount = 1024
	fullScan := flag.Bool("full", false, "run a full 65535 port scan")
	flag.StringVar(&targetip, "ip", "0.0.0.0", "target ip address")
	flag.Parse()

	if *fullScan {
		targetPortCount = 65535
	}

	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= targetPortCount; i++ {
			ports <- i
		}
	}()

	for i := 0; i < targetPortCount; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d,", port)
	}
	fmt.Printf("\n")
}
