

# Goals
- mapping the requested flows
- mapping the actual flows
- presenting a graphic that materialize flows

# Target
- extension of the reminder apps ?

# Execution mode

In each mode below, binary is the same, only `mode` param value change.

## Agent

"One shot mode", agent start, analyse connections for the host on which it is running and sends 
results to `apiserver` module.
Agent mode list:
- check-agent: create pseudo servers and test connections
- probe-agent: lists the ports actually listening

## Server
Long running mode : receives report from agents, persist them into database and
render graph. 
Server mode list : 
- apiserver: receives report from agents, persist them into database
- graphviz: presents dashboard graphics (d3js)

# Params list

```
usage: 
    portchecker apiserver
      or
    portchecker graphviz
      or
    portchecker check-agent --mapping_file_url<mappingFileUrl> --analysis_id <analysis_id>
      or
    portchecker probe-agent --analysis_id <analysis_id>
```

__Mode__: run mode. Must be one of : 
* check-agent
* probe-agent
* apiserver
* graphviz

__mappingFileUrl__ : local path to file that contains the description of fluxes to tests

__analysis_id__: ID that identify the analysis number (agents have no valid solution to determine which is the current test session ID)

# Requirements
- database

=> For now, provisionned with docker-compose.yml file : only for dev mode

# Global launch mode

* Start apiserver
* Launch agents on each instance to test : The agents can be launched via Ansible, create their report and send to the server API. Max. Execution = 2 min
* Launch graphviz to generate static web server that contains files to generate D3JS reports. Datas in graph have been extracted by apiserver


