package conf

import (
	"time"
	"os"
	"fmt"
	"strconv"
)

type Config struct {

	Mode string `arg:"positional"`
	MappingFileUrl string `help:"Url to network config file to check"`
	AnalysisId int `help:"Analysis ID to use when sending result on API server"`

	ConfigFilePath string
	Verbose bool
	ApiServerUrl string
	Hostname string

	// Probe Agent Config
	ProbeConcurrentThreadsNumber int
	ProbeTimeoutDuration         time.Duration
	ProbePortRangeStart          int
	ProbePortRangeStop           int

	// Check Agent Config
	Timeout int

	// API server config
	ApiServerListenPort int
}

func NewConfig() Config{
	res:= Config{}

	res.ApiServerUrl = os.Getenv("APISERVER_URL")

	if res.ApiServerUrl == "" {
		fmt.Fprintf(os.Stdout, "WARN: APISERVER_URL env var is missing, using default localhost one.\n")
		res.ApiServerUrl = "http://localhost:8090"
	}
	res.Hostname, _ = os.Hostname()

	res.ProbeConcurrentThreadsNumber = ParseIntEnvVar("PROBE_AGENT_CONCURRENT_THREADS_NUMBER", PROBE_AGENT_CONCURRENT_THREADS_NUMBER)
	res.ProbeTimeoutDuration = ParseDurationEnvVar("PROBE_AGENT_TIMEOUT_DURATION", PROBE_AGENT_TIMEOUT_DURATION)
	res.ProbePortRangeStart = ParseIntEnvVar("PROBE_AGENT_CHECK_PORT_RANGE_START", PROBE_AGENT_CHECK_PORT_RANGE_START)
	res.ProbePortRangeStop = ParseIntEnvVar("PROBE_AGENT_CHECK_PORT_RANGE_STOP", PROBE_AGENT_CHECK_PORT_RANGE_STOP)
	res.Timeout = ParseIntEnvVar("CHECK_AGENT_TIMEOUT", CHECK_AGENT_TIMEOUT)

	res.ApiServerListenPort = ParseIntEnvVar("APISERVER_LISTEN_PORT", APISERVER_LISTEN_PORT)

	return res
}

func ParseIntEnvVar(name string, defaultValue int) int {
	val, exist := os.LookupEnv(name)
	if ! exist {
		return defaultValue
	} else {
		if val, err := strconv.Atoi(val); err == nil {
			return val
		} else {
			fmt.Fprintf(os.Stderr, fmt.Sprintf("%v env var must be an int and cannot be parsed as int. Using default value %v", name, defaultValue))
			return defaultValue
		}
	}
}
func ParseDurationEnvVar(name string, defaultValue time.Duration) time.Duration {
	val, exist := os.LookupEnv(name)
	if ! exist {
		return defaultValue
	} else {
		if val, err := strconv.Atoi(val); err == nil {
			return time.Duration(val)
		} else {
			fmt.Fprintf(os.Stderr, fmt.Sprintf("%v env var must be an time.Duration and cannot be parsed as int. Using default value %v", name, defaultValue))
			return defaultValue
		}
	}
}
