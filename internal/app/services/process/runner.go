package process

import (
    "fmt"
    "os/exec"
    "log"
    "time"
)

type Service struct {
    ServiceName string
    Namespace string
    FromPort int
    ToPort int
    StartAt time.Time
    FailedCount int
}

type ServiceMonitor struct {
    Services []*Service
}

var _serviceMonitor = ServiceMonitor{
    Services: []*Service{},
}

func registerService(namespace string, serviceName string, fromPort int, toPort int) bool {

    service := getService(namespace, serviceName)
    if service != nil {
        return false
    }

    service = &Service{
        ServiceName:serviceName,
        Namespace: namespace,
        FromPort: fromPort,
        ToPort: toPort,
        StartAt: time.Now(),
        FailedCount: 0,
    }
    _serviceMonitor.Services = append(_serviceMonitor.Services, service)

    go runService(service)
    return true
}

func runService(service *Service) {
    cmd := exec.Command("kubectl", "port-forward", "svc/" + service.ServiceName , "-n", service.Namespace, fmt.Sprintf("%d:%d", service.ToPort, service.FromPort))
    err := cmd.Start()
    if err != nil {
       log.Fatal(err)
    }
    log.Printf("Waiting for command to finish...")
    err = cmd.Wait()
    if err != nil {
        log.Println(err)
        service.FailedCount ++
        if service.FailedCount < 5 {
            go runService(service)
        }
    }
}

func getService(namespace string, serviceName string) *Service {
    for _, service := range _serviceMonitor.Services {
        if service.ServiceName == serviceName && service.Namespace == namespace {
            return service
        }
    }
    return nil
}

func PortForward(namespace string, serviceName string, fromPort int, toPort int) bool {
    success := registerService(namespace, serviceName, fromPort, toPort)
    if !success {
        return false
    }

    return true

}