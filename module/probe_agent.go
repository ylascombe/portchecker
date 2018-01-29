package module

import (
	"github.com/anvie/port-scanner"
	"fmt"
	"portchecker/db_models"
	"encoding/json"
	"strings"
	"portchecker/services"
	"portchecker/conf"
	"time"
)

func ProbeAgent(config conf.Config) {

	fmt.Println(config)
	postUrl := fmt.Sprintf("%v/v1/hostname/%v/analysis_id/%v/probe_agent/", config.Hostname, config.AnalysisId)

	if config.Verbose {
		fmt.Printf("scanning port %d-%d...\n", config.ProbePortRangeStart, config.ProbePortRangeStop)
	}

	openedPorts := FindOpenedPort(
		config.ProbePortRangeStart,
		config.ProbePortRangeStop,
		config.ProbeTimeoutDuration,
		config.ProbeConcurrentThreadsNumber,
	)

	probeAgent := db_models.ProbeAgent{
		Hostname:          config.Hostname,
		AnalysisId:        config.AnalysisId,
		OpenedPorts:       openedPorts,
		OpenedPortsString: strings.Trim(strings.Join(strings.Fields(fmt.Sprint(openedPorts)), ","), "[]"),
	}

	res, _ := json.Marshal(probeAgent)

	services.SendResultToApiserver(postUrl, "probe-agent", res)
}

func FindOpenedPort(portStart int, portEnds int, timeoutDuration time.Duration, concurrentThreadsNumber int) []int {

	// scan localhost with a 2 second timeout per port in 5 concurrent threads
	ps := portscanner.NewPortScanner("localhost", timeoutDuration, concurrentThreadsNumber )

	// get opened port
	openedPorts := ps.GetOpenedPort(portStart, portEnds)

	return openedPorts
}
