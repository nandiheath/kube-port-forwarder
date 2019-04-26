package app

import (
    "flag"
    "fmt"
    "github.com/kubernetes/client-go/kubernetes"
    "github.com/kubernetes/client-go/tools/clientcmd"
    v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
    //"k8s.io/apimachinery/pkg/apis/meta/v1"
    "log"
    "os/user"
    "path/filepath"
)

func LoadKube() {

    // Path the kube config
    usr, err := user.Current()
    if err != nil {
        log.Fatal( err )
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
        panic(err.Error())
    }

    // creates the clientset
    clientset, err := kubernetes.NewForConfig(config)

    if err != nil {
        log.Fatal(err)
        panic(err)
    }

    // access the API to list pods
    pods, _:= clientset.CoreV1().Namespaces().List(v1.ListOptions{})

    fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
}