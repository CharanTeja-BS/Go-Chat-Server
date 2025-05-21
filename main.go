package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	// Command line flags
	target := flag.String("target", "scanme.nmap.org", "Target IP or domain to scan")
	startPort := flag.Int("start", 1, "Start port number")
	endPort := flag.Int("end", 1024, "End port number")
	flag.Parse()

	fmt.Printf("🔍 Scanning %s from port %d to %d\n", *target, *startPort, *endPort)

	var wg sync.WaitGroup
	portsChan := make(chan int, 100) // buffer for ports to scan
	results := make(chan int)         // open ports

	// Worker goroutine function
	worker := func() {
		for port := range portsChan {
			address := fmt.Sprintf("%s:%d", *target, port)
			conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
			if err == nil {
				conn.Close()
				results <- port
			} else {
				results <- 0 // indicate closed port by 0 (will ignore later)
			}
			wg.Done()
		}
	}

	// Start 100 workers (tune as needed)
	for i := 0; i < 100; i++ {
		go worker()
	}

	// Add all ports to waitgroup and queue
	for port := *startPort; port <= *endPort; port++ {
		wg.Add(1)
		portsChan <- port
	}
	close(portsChan) // no more ports incoming

	// Collect open ports concurrently
	go func() {
		wg.Wait()
		close(results)
	}()

	// Print open ports
	for port := range results {
		if port != 0 {
			fmt.Printf("✅ Port %d is OPEN\n", port)
		}
	}

	fmt.Println("✅ Scan complete.")
}
