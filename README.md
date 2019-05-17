# Kube Port Forwarder

This is a web UI for managing `kubectl port-forward` commands.
With the help of this tool, user can port-forward services on kubernetes cluster with different namespace easily.

## Dependencies

- `kubectl`
- `~/.kube/config` file exists
- access rights to the current kubernetes cluster

### Attention

- Multi-cluster not supported
  - Work only with the current cluster context from your `~/.kube/config`
- Location of the kube config file
  - fixed at `~/.kube/config`
- One service only be able to forward one port
- All port-forwards are spawned as child process
  - will be killed directly if the server process died

## Start

```bash

git clone git@github.com:nandiheath/kube-port-forwarder.git

cd kube-port-forwarder

# Start the service at port 8080
./scripts/server.sh start

# or start the service at custom port
./scripts/server.sh start 9000

# Stop the service
./scripts/server.sh stop

```