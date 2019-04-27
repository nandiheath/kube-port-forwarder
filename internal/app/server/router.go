package server

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/nandiheath/kube-port-forwarder/internal/app/services/k8s"
	"github.com/nandiheath/kube-port-forwarder/internal/app/services/process"
	"net/http"
	"path/filepath"
	"strconv"
)

var templatesDir = "web/views"

func SetupHTMLRender(r *gin.Engine) {
	renderer := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.tpl")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/pages/*.tpl")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		renderer.AddFromFiles(filepath.Base(include), files...)
	}

	r.HTMLRender = renderer
}

func Route(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		namespaces := k8s.GetNamespaces()

		c.HTML(200, "namespace.tpl", gin.H{
			"title":      "Welcome!",
			"namespaces": namespaces,
		})
	})

	r.GET("/api/namespaces", func(c *gin.Context) {
		namespaces := k8s.GetNamespaces()
		c.HTML(http.StatusOK, "namespace.tmpl", gin.H{
			"title":      "Namespace",
			"namespaces": namespaces,
		})
	})

	r.GET("/forward", func(c *gin.Context) {
		namespace := c.Query("namespace")
		serviceName := c.Query("serviceName")
		from, err := strconv.Atoi(c.Query("from"))
		to, err := strconv.Atoi(c.Query("to"))

		if err != nil || !process.PortForward(namespace, serviceName, from, to) {
			c.JSON(400, gin.H{})
		} else {
			c.JSON(200, gin.H{
				"success": true,
			})
		}

	})
}
