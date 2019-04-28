package server

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/nandiheath/kube-port-forwarder/internal/app/services/k8s"
	"github.com/nandiheath/kube-port-forwarder/internal/app/services/process"
	"html/template"
	v12 "k8s.io/api/core/v1"
	"net/http"
	"path/filepath"
)

// Binding for PUT /namespace/:namespace
type PortForward struct {
	ServiceName string `form:"service" json:"service" binding:"required"`
	FromPort    int    `form:"from_port" json:"from_port" binding:"required"`
	ToPort      int    `form:"to_port" json:"to_port" binding:"required"`
}

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
		renderer.AddFromFilesFuncs(filepath.Base(include), template.FuncMap{
			"check": check,
		}, files...)
	}

	r.HTMLRender = renderer
}

func Route(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		namespaces := k8s.GetNamespaces()
		c.HTML(200, "namespaces.tpl", gin.H{
			"title":      "Welcome!",
			"namespaces": namespaces,
		})
	})

	r.GET("/namespace/:namespace", func(c *gin.Context) {
		namespace := c.Param("namespace")
		showNamespacePageWithError(c, namespace, "")
	})

	r.PUT("/namespace/:namespace", func(c *gin.Context) {
		namespace := c.Param("namespace")
		var json PortForward
		if err := c.ShouldBindJSON(&json); err != nil {
			showNamespacePageWithError(c, namespace, "invalid parameter")
			return
		}

		// TODO: show the type of error
		if !process.PortForward(namespace, json.ServiceName, json.FromPort, json.ToPort) {
			showNamespacePageWithError(c, namespace, "cannot forward port")
			return
		}

		showNamespacePageWithError(c, namespace, "")

	})
}

func showNamespacePageWithError(c *gin.Context, namespace string, error string) {
	services := k8s.GetServices(namespace)
	c.HTML(http.StatusOK, "namespace.tpl", gin.H{
		"title":    namespace,
		"services": services,
		"error":    error,
	})
}

func check(service *v12.Service) string {
	return service.Name
}
