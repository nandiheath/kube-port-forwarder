package main

import (
	"fmt"
	"github.com/nandiheath/kube-port-forwarder/internal/app/server"
	"io/ioutil"
	"os"
)

func main() {
	d1 := []byte(fmt.Sprintf("%d", os.Getpid()))
	err := ioutil.WriteFile("/Users/nandi0315/workspace/kube-port-forwarder/run.pid", d1, 0644)
	if err != nil {
		panic(err)
	}
	server.StartServer()

}
