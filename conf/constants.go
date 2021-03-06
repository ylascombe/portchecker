package conf

import "time"


const PROBE_AGENT_CONCURRENT_THREADS_NUMBER = 200
const PROBE_AGENT_TIMEOUT_DURATION = 100 * time.Millisecond
const PROBE_AGENT_CHECK_PORT_RANGE_START = 1
const PROBE_AGENT_CHECK_PORT_RANGE_STOP = 15000

const CHECK_AGENT_TIMEOUT = 20

const APISERVER_LISTEN_PORT = 8090
