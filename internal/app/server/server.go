package server

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/nandiheath/kube-port-forwarder/internal/app/services/k8s"
)

func StartServer() {
	k8s.KubeInit()
	r := gin.Default()

	// serve static files
	r.Use(static.Serve("/static", static.LocalFile("./web/static", false)))

	SetupHTMLRender(r)
	Route(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
