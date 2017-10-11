package main

import "fmt"
import "flag"
import (
	"os"
	"portchecker/services"
	"portchecker/utils"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: portchecker [mode in [check-agent|probe-agent|apiserver|graphviz] ]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

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
		services.DoWork(*config, hostname, 20)
	case "probe-agent":
		fmt.Fprintf(os.Stderr, "Not implemented\n")
	case "apiserver":
		fmt.Fprintf(os.Stderr, "Not implemented\n")
	case "graphviz":
		fmt.Fprintf(os.Stderr, "Not implemented\n")

	default:
		fmt.Fprintf(os.Stderr, "Invalid mode %s\n", mode)
		os.Exit(1)
	}


	os.Exit(0)
}
