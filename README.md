

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

# Requirements
- database

The agents are launched via Ansible, create their report and send to the server API. Max. Execution = 2 min
