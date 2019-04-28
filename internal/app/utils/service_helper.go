package utils

import (
	"github.com/nandiheath/kube-port-forwarder/internal/app/services/process"
	v12 "k8s.io/api/core/v1"
)

type ServicePort struct {
	Name       string
	Port       int
	TargetPort int
}

type Service struct {
	Name          string
	Type          string
	Ports         []ServicePort
	Status        string
	Forwarded     bool
	ForwardedFrom int
	ForwardedTo   int
}

func MapForwardedServiceAndService(kServices []v12.Service, forwardedServices []*process.ForwardedService) []Service {
	var serviceList []Service
	for _, kService := range kServices {
		service := Service{
			Name:      kService.Name,
			Type:      string(kService.Spec.Type),
			Ports:     []ServicePort{},
			Forwarded: false,
			Status:    "N/A",
		}

		var forwardedService *process.ForwardedService
		for _, fService := range forwardedServices {
			if fService.ServiceName == kService.Name {
				forwardedService = fService
				service.Status = forwardedService.Status
				service.ForwardedFrom = forwardedService.FromPort
				service.ForwardedTo = forwardedService.ToPort
				service.Forwarded = true
				break
			}
		}

		for _, port := range kService.Spec.Ports {

			servicePort := ServicePort{
				Name:       port.Name,
				Port:       int(port.Port),
				TargetPort: port.TargetPort.IntValue(),
			}

			service.Ports = append(service.Ports, servicePort)
		}
		serviceList = append(serviceList, service)
	}
	return serviceList
}
