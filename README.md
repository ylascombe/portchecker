

# Goals
- mapping the requested flows
- mapping the actual flows
- presenting a graphic that materialize flows

# Target
- extension of the reminder apps ?

# Execution mode
- check: create pseudo servers and test connections
- probe: lists the ports actually listening
- apiserver: interface between agents and clients: manages persistence
- graphviz: presents dashboard graphics (d3js)

In each mode below, binary is the same, only `mode` param value change.

# Requirements
- database

The agents are launched via Ansible, create their report and send to the server API. Max. Execution = 2 min
