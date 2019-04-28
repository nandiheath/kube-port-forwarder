# Kube Port Forwarder

This is a web UI for managing `kubectl port-forward` commands.
With the help of this tool, user can port-forward services on kubernetes cluster with different namespace easily.

## Dependencies

- `kubectl`
- `~/.kube/config` file exists
- access rights to the current kubernetes cluster

### Attensions

- Multi-cluster not supported
  - Work only with the current cluster context from your `~/.kube/config`
- Location of the kube config file
  - support only `~/.kube/config`
- One service only be able to forward one port
- All port-forwards are spawned as child process
  - will be killed directly if the go service died

## Start

```bash

go build

# Run as default 8080 port
go run cmd/port_forwarder/main.go

# OR run with custom port
PORT=8081 go run cmd/port_forwarder/main.go
```

This will create the webUI at `localhost:8080`