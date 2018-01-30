package main

import (
	"os"
	"portchecker/utils"
	"portchecker/module"
	"fmt"
	"flag"
	"github.com/alexflint/go-arg"
	"portchecker/conf"
	"errors"
)



func usage() {
	fmt.Fprintf(os.Stderr, "usage: \n" +
		"\tportchecker apiserver\n" +
		"or\n" +
		"\tportchecker graphviz\n" +
		"or\n" +
		"\tportchecker check-agent --mapping_file_url<mappingFileUrl> --analysisid <analysisid>\n" +
		"or\n" +
		"\tportchecker probe-agent --analysisid <analysisid>\n\nParameters desc:\n")
	flag.PrintDefaults()
	os.Exit(2)
}
//
//func init() {
//	config := conf.NewConfig()
//	config.Mode = *flag.String("mode", "apiserver", "portchecker mode in [apiserver|graphviz|check-agent|probe-agent]")
//	config.AnalysisId = *flag.Int( "analysisid", 1, "unique anlysis id")
//	config.ConfigFilePath = *flag.String( "mapping_file_url", "", "Local path to mapping description file")
//	config.Verbose = *flag.Bool("verbose", false, "Log verbose mode")
//}

func main() {
	config := conf.NewConfig()
	arg.MustParse(&config)


	 flag.Usage = usage
	flag.Parse()

	paramErrors := checkParam(config)

	if paramErrors != nil {
		fmt.Fprintf(os.Stderr, "Invalid parameters %s\n", paramErrors)
		os.Exit(1)
	}

	if config.Verbose {
		fmt.Println(fmt.Sprintf("Ask to start in %v mode", config.Mode))
	}

	switch config.Mode {
	case "check-agent":
		networkConfig, err := utils.UnmarshallFromFile(config.ConfigFilePath)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot load config from given path : %v. Err : %v", config.ConfigFilePath, err)
			os.Exit(1)
		}

		checkResult, err := module.StartCheckAgent(*networkConfig, config)

		if err != nil {
			fmt.Fprintf(os.Stderr, "An error occurred %v", err)
		}

		err = module.ProcessCheckAgentResult(*networkConfig, *checkResult, config)

	case "probe-agent":
		module.ProbeAgent(config)
	case "apiserver":
		module.StartApiServer(config)

	case "graphviz":
		fmt.Fprintf(os.Stderr, "Not implemented\n")

	default:
		fmt.Fprintf(os.Stderr, "Invalid mode %s\n", config.Mode)
		os.Exit(1)
	}

	os.Exit(0)
}

func checkParam(config conf.Config) error {
	var required []string
	switch config.Mode {
	case "check-agent":
		required = []string{"mapping_file_url", "analysisid"}
	case "probe-agent":
		required = []string{"analysisid"}
	case "apiserver":
		required = []string{}
	case "graphviz":
		required = []string{}
	default:
		required = []string{}
	}

	flag.Parse()

	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })

	if len(required) > 0 {
		return nil
	}

	missings := ""
	for _, req := range required {
		if !seen[req] {
			// or possibly use `log.Fatalf` instead of:
			msg := fmt.Sprintf("ERROR: missing required -%s argument/flag\n\n", req)
			missings += "\n" + msg
		}
	}

	if len(missings) > 0 {
		return errors.New(missings)
	} else {
		return nil
	}

}
