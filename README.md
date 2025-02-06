# ogrex WORK IN PROGRESS
Ogrex is simple reversed proxy project written in Go.

# Current features
- Reversed proxy
- Load balancing

# How to use

Currently availiable commands: 

`run path/to/config.yaml` runs proxy with given configuration


# Configuration
Configuration has to be specified in .yaml file and passed after `run` command:

```
.yaml configuration file example

server:
 port: 8080

someService:
 url: '/someService'
 services:
  - https://subservice-1-url.com/
  - https://subservice-2-url.com

someOtherService:
 url: '/someService2'
 services:
  - https://example.com/

```