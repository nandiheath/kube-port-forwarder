package process

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

type ForwardedService struct {
	ServiceName string
	Namespace   string
	FromPort    int
	ToPort      int
	StartAt     time.Time
	FailedCount int
	Status      string
}

type ForwardedServiceList struct {
	Services []*ForwardedService
}

var _serviceMonitor = ForwardedServiceList{
	Services: []*ForwardedService{},
}

func registerService(namespace string, serviceName string, fromPort int, toPort int) bool {

	service := getService(namespace, serviceName)
	if service != nil && service.Status != "Failed" {
		return false
	}

	if service == nil {
		service = &ForwardedService{
			ServiceName: serviceName,
			Namespace:   namespace,
			FromPort:    fromPort,
			ToPort:      toPort,
			StartAt:     time.Now(),
			FailedCount: 0,
			Status:      "Pending",
		}
		_serviceMonitor.Services = append(_serviceMonitor.Services, service)
	} else {
		service.FromPort = fromPort
		service.ToPort = toPort
		service.FailedCount = 0
		service.StartAt = time.Now()
	}

	go runService(service)
	return true
}

func runService(service *ForwardedService) {
	cmd := exec.Command("kubectl", "port-forward", "svc/"+service.ServiceName, "-n", service.Namespace, fmt.Sprintf("%d:%d", service.ToPort, service.FromPort))
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	service.Status = "Forwarded"
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	if err != nil {
		log.Println(err)
		service.FailedCount++
		if service.FailedCount < 5 {
			go runService(service)
		} else {
			service.Status = "Failed"
		}
	}
}

func getService(namespace string, serviceName string) *ForwardedService {
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

func GetForwardedService(namespace string) []*ForwardedService {
	var services []*ForwardedService
	for _, service := range _serviceMonitor.Services {
		if service.Namespace == namespace {
			services = append(services, service)
		}
	}
	return services
}
