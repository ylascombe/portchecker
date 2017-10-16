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
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: portchecker [mode in [check-agent|probe-agent|apiserver|graphviz] ]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	serverUrl := "http://localhost:8090/v1/check_agents/10"
	args := flag.Args()
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Portchecker mode or config file path are missing .")
		os.Exit(1)
	}

	mode := args[0]
	configFilePath := args[1]
	fmt.Println("Starting in", mode, " mode")
	config, err := utils.UnmarshallFromFile(configFilePath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot load config from given path : %v. Err : %v", configFilePath, err)
		os.Exit(1)
	}
	switch mode {
	case "check-agent":
		hostname, _ := os.Hostname()
		checkResult, err := module.StartCheckAgent(*config, hostname, 20)

		if err != nil {
			fmt.Fprintf(os.Stderr, "An error occurred %v", err)
		}

		checkAgent, err := module.ProcessCheckAgentResult(*config, *checkResult, hostname)
		fmt.Println("")
		fmt.Println("")
		res, _ := json.Marshal(checkAgent)

		postRes, err := http.Post(serverUrl, "application/json", bytes.NewBuffer(res))
		fmt.Fprintf(os.Stdout, "POST Result \n%v. Err %v", postRes, err)

		fmt.Println("")
		fmt.Println("")
		fmt.Fprintf(os.Stdout, "JSON RESULT \n%v", string(res))


	case "probe-agent":
		fmt.Fprintf(os.Stderr, "Not implemented\n")
	case "apiserver":
		fmt.Fprintf(os.Stderr, "Not implemented\n")
		module.StartApiServer()

	case "graphviz":
		fmt.Fprintf(os.Stderr, "Not implemented\n")

	default:
		fmt.Fprintf(os.Stderr, "Invalid mode %s\n", mode)
		os.Exit(1)
	}


	os.Exit(0)
}
