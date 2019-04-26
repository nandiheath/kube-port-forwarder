package k8s

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
    "os/user"
)

type KubeContext struct {
    Context struct{
        Cluster string `yaml:"cluster"`
        Namespace string `yaml:"namespace"`
        User string `yaml:"user"`
    }
    Name string
}

type KubeUser struct {
    Name string
    User struct{
        AuthProvider struct{
            Config struct{
                AccessToken string `yaml:"access-token"`
            } `yaml:"config"`
            Name string
        } `yaml:"auth-provider"`
    }

}

type KubeConfig struct {
    ApiVersion string `yaml:"apiVersion"`
    Contexts []KubeContext `yaml:"contexts,flow"`
    CurrentConext string `yaml:"current-context"`
    Users []KubeUser `yaml:"users,flow"`
}

var KubeAuthenication string

func LoadKubeFile() {
    usr, err := user.Current()
    if err != nil {
        log.Fatal( err )
    }

    data, err := ioutil.ReadFile(usr.HomeDir + "/.kube/config")
    if err != nil {
        panic(err)
    }
    parsedConfig := KubeConfig{}
    err = yaml.Unmarshal(data, &parsedConfig)
    if err != nil {
        log.Fatal(err)
    }
    //fmt.Println(parsedConfig.CurrentConext)
    parseKubeConfig(parsedConfig, &KubeAuthenication)

    fmt.Println(KubeAuthenication)
}

func parseKubeConfig(config KubeConfig, accessToken *string) {

    currentUserName := ""
    currentContextName := config.CurrentConext

    for _, context := range config.Contexts {
        if context.Name == currentContextName {
            currentUserName = context.Context.User
            break
        }
    }

    if currentUserName == "" {
        panic("Cannot find the authenication within the kube config")
    }

    for _, user := range config.Users {
        if user.Name == currentUserName {
            *accessToken = user.User.AuthProvider.Config.AccessToken
            return
        }
    }
}