package server

import (
    "github.com/gin-gonic/gin"
    "github.com/nandiheath/kube-port-forwarder/internal/app/services/k8s"
    "github.com/nandiheath/kube-port-forwarder/internal/app/services/process"
    "strconv"
)

func Route(r *gin.Engine) {
    r.GET("/api/namespaces", func(c *gin.Context) {
        namespaces := k8s.GetNamespaces()
        c.JSON(200, gin.H{
            "message": namespaces,
        })
    })

    r.GET("/forward", func(c *gin.Context) {
        namespace := c.Query("namespace")
        serviceName := c.Query("serviceName")
        from, err := strconv.Atoi(c.Query("from"))
        to, err := strconv.Atoi(c.Query("to"))

        if err != nil || !process.PortForward(namespace, serviceName, from, to) {
            c.JSON(400, gin.H{
            })
        } else {
            c.JSON(200, gin.H{
                "success": true,
            })
        }


    })
}