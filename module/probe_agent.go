package module

import (
	"time"
	"github.com/anvie/port-scanner"
	"fmt"
	"portchecker/db_models"
	"bytes"
	"os"
	"net/http"
	"encoding/json"
	"strings"
)

const CONCURRENT_THREADS_NUMBER = 200
const TIMEOUT_DURATION = 100 * time.Millisecond

func ProbeAgent(hostname string, analysisId int, postUrl string) {
	openedPorts := FindOpenedPort(1, 15000)

	probeAgent := db_models.ProbeAgent{
		Hostname: hostname,
		AnalysisId: analysisId,
		OpenedPorts: openedPorts,
		OpenedPortsString: strings.Trim(strings.Join(strings.Fields(fmt.Sprint(openedPorts)), ","), "[]"),
	}

	res, _ := json.Marshal(probeAgent)

	fmt.Println("")
	fmt.Println("")
	fmt.Println(string(res))
	postRes, err := http.Post(postUrl, "application/json", bytes.NewBuffer(res))
	fmt.Fprintf(os.Stdout, "POST Result \n%v. Err %v", postRes, err)

}

func FindOpenedPort(portStart int, portEnds int) []int {
	// scan localhost with a 2 second timeout per port in 5 concurrent threads
	ps := portscanner.NewPortScanner("localhost", TIMEOUT_DURATION, CONCURRENT_THREADS_NUMBER)


	// get opened port
	fmt.Printf("scanning port %d-%d...\n", portStart, portEnds)

	openedPorts := ps.GetOpenedPort(portStart, portEnds)


	// for i := 0; i < len(openedPorts); i++ {
 		// port := openedPorts[i]
		// fmt.Print(" ", port, " [open]")
		// fmt.Println("  -->  ", ps.DescribePort(port))
	// }

	return openedPorts
}
