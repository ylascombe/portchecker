package main

import (
	"os"
	"portchecker/utils"
	"portchecker/module"
	"fmt"
	"flag"
)

var mode string;
var configFilePath string
var analysisId int
var quiet bool

func usage() {
	fmt.Fprintf(os.Stderr, "usage: \n" +
		"\tportchecker\n" +
		"or\n" +
		"\tportchecker apiserver\n" +
		"or\n" +
		"\tportchecker graphviz\n" +
		"or\n" +
		"\tportchecker check-agent --mapping_file_url<mappingFileUrl> --analysis_id <analysis_id>\n" +
		"or\n" +
		"\tportchecker probe-agent --analysis_id <analysis_id>\n\nParameters desc:\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	flag.StringVar(&mode, "mode", "apiserver", "portchecker mode in [apiserver|graphviz|check-agent|probe-agent]")
	flag.IntVar(&analysisId, "analysis_id", 1, "unique anlysis id")
	flag.StringVar(&configFilePath, "mapping_file_url", "", "Local path to mapping description file")
	flag.BoolVar(&quiet, "quiet", false, "Log quiet mode")
}

func main() {
	flag.Usage = usage
	flag.Parse()

	checkParam()

	apiServerUrl := os.Getenv("APISERVER_URL")

	if apiServerUrl == "" {
		fmt.Fprintf(os.Stdout, "WARN: APISERVER_URL env var is missing, using default localhost one.\n")
		apiServerUrl = "http://localhost:8090"
	}
	hostname, _ := os.Hostname()

	fmt.Println(fmt.Sprintf("Ask to start in %v mode", mode))

	switch mode {
	case "check-agent":
		config, err := utils.UnmarshallFromFile(configFilePath)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot load config from given path : %v. Err : %v", configFilePath, err)
			os.Exit(1)
		}

		checkResult, err := module.StartCheckAgent(*config, hostname, 20)

		if err != nil {
			fmt.Fprintf(os.Stderr, "An error occurred %v", err)
		}

		finalUrl := fmt.Sprintf("%v/v1/hostname/%v/analysis_id/%v/check_agents/", apiServerUrl, hostname, analysisId)
		err = module.ProcessCheckAgentResult(*config, *checkResult, hostname, finalUrl)

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

func checkParam() {
	var required []string
	switch mode {
	case "check-agent":
		required = []string{"mode", "mapping_file_url", "analysis_id"}
	case "probe-agent":
		required = []string{"mode", "analysis_id"}
	case "apiserver":
		required = []string{"mode"}
	case "graphviz":
		required = []string{"mode"}
	default:
		required = []string{"mode"}
	}

	flag.Parse()

	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range required {
		if !seen[req] {
			// or possibly use `log.Fatalf` instead of:
			fmt.Fprintf(os.Stderr, "ERROR: missing required -%s argument/flag\n\n", req)
			usage()
		}
	}
}
