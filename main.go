package main

import (
	"os"
	"portchecker/utils"
	"portchecker/module"
	"fmt"
	"flag"
	"encoding/json"
	"net/http"
	"bytes"
	"strconv"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: portchecker <mode in [check-agent|probe-agent|apiserver|graphviz]> <mappingFileUrl> <analysis_id>\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	apiServerUrl := os.Getenv("APISERVER_URL")

	if apiServerUrl == "" {
		fmt.Fprintf(os.Stdout, "WARN: APISERVER_URL env var is missing, using default localhost one.\n")
		apiServerUrl = "http://localhost:8090"
	}
	hostname, _ := os.Hostname()

	args := flag.Args()
	if len(args) < 3 {
		usage()
		os.Exit(1)
	}

	mode := args[0]
	configFilePath := args[1]
	analysisIdStr := args[2]
	analysisId, err := strconv.Atoi(analysisIdStr)



	if err != nil {
		fmt.Fprintf(os.Stderr, "analysis_id parameter is invalid. %v", err)
		os.Exit(3)
	}

	fmt.Println("Ask to start in", mode, " mode")
	config, err := utils.UnmarshallFromFile(configFilePath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot load config from given path : %v. Err : %v", configFilePath, err)
		os.Exit(1)
	}
	switch mode {
	case "check-agent":
		checkResult, err := module.StartCheckAgent(*config, hostname, 20)

		if err != nil {
			fmt.Fprintf(os.Stderr, "An error occurred %v", err)
		}

		checkAgent, err := module.ProcessCheckAgentResult(*config, *checkResult, hostname)
		fmt.Println("")
		fmt.Println("")
		res, _ := json.Marshal(checkAgent)


		finalUrl := fmt.Sprintf("%v/v1/hostname/%v/analysis_id/%v/check_agents/", apiServerUrl, hostname, analysisId)
		postRes, err := http.Post(finalUrl, "application/json", bytes.NewBuffer(res))
		fmt.Fprintf(os.Stdout, "POST Result \n%v. Err %v", postRes, err)

		fmt.Println("")
		fmt.Println("")
		fmt.Fprintf(os.Stdout, "JSON RESULT \n%v", string(res))


	case "probe-agent":
		finalUrl := fmt.Sprintf("%v/v1/hostname/%v/analysis_id/%v/probe_agent/", apiServerUrl, hostname, analysisId)
		module.ProbeAgent(hostname, analysisId, finalUrl)

	case "apiserver":
		module.StartApiServer()

	case "graphviz":
		fmt.Fprintf(os.Stderr, "Not implemented\n")

	default:
		fmt.Fprintf(os.Stderr, "Invalid mode %s\n", mode)
		os.Exit(1)
	}


	os.Exit(0)
}
