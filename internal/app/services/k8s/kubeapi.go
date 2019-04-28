package k8s

import (
	"flag"
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/clientcmd"
	//"k8s.io/apimachinery/pkg/apis/meta/v1"
	"os/user"
	"path/filepath"
)

var _client *kubernetes.Clientset

func getKubeClient() *kubernetes.Clientset {
	if _client != nil {
		return _client
	}

	// Path the kube config
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	var kubeconfig *string
	if home := usr.HomeDir; home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	// creates the clientset
	_client, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	return _client
}

func KubeInit() {
	getKubeClient()
}

func GetNamespaces() []v12.Namespace {
	client := getKubeClient()

	// access the API to list pods
	namespaces, _ := client.CoreV1().Namespaces().List(v1.ListOptions{})

	return namespaces.Items
}

func GetServices(namespace string) []v12.Service {
	client := getKubeClient()

	// access the API to list pods
	services, _ := client.CoreV1().Services(namespace).List(v1.ListOptions{})

	return services.Items
}
