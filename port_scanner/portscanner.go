package port_scanner

import (
	"fmt"
	"time"
	"github.com/anvie/port-scanner"
)

const CONCURRENT_THREADS_NUMBER = 200
const TIMEOUT_DURATION = 100 * time.Millisecond

func FindOpenedPort(portStart int, portEnds int){
	// scan localhost with a 2 second timeout per port in 5 concurrent threads
	ps := portscanner.NewPortScanner("localhost", TIMEOUT_DURATION, CONCURRENT_THREADS_NUMBER)

	// get opened port
	fmt.Printf("scanning port %d-%d...\n", portStart, portEnds)

	openedPorts := ps.GetOpenedPort(portStart, portEnds)

	for i := 0; i < len(openedPorts); i++ {
		port := openedPorts[i]
		fmt.Print(" ", port, " [open]")
		fmt.Println("  -->  ", ps.DescribePort(port))
	}
}
