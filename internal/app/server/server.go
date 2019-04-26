package server

import (
    "github.com/gin-gonic/gin"
    "github.com/nandiheath/kube-port-forwarder/internal/app/services/k8s"
)

func StartServer() {
    k8s.KubeInit()
    r := gin.Default()

    Route(r)
    r.Run() // listen and serve on 0.0.0.0:8080
}