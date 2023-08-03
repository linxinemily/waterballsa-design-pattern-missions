package domain

import (
	"fmt"
)

type LoadBalancing struct {
	next      HttpClient
	currentId int
}

func NewLoadBalancing(next HttpClient) *LoadBalancing {
	return &LoadBalancing{next: next}
}

func (l *LoadBalancing) SendRequest(req *HttpRequest) (*HttpResponse, error) {

	targetIp, err := l.determineTargetIP(req)
	if err != nil {
		return nil, err
	}

	req.TargetIp = targetIp
	fmt.Println("[Load Balancing] target ip:", targetIp)

	return l.next.SendRequest(req)
}

func (l *LoadBalancing) determineTargetIP(req *HttpRequest) (string, error) {
	if len(req.IPs) == 0 {
		return req.Host, nil
	}

	ip, err := l.findNextActiveIP(req.IPs)
	if err != nil {
		return "", err
	}

	return ip, nil
}

func (l *LoadBalancing) findNextActiveIP(ips []IP) (string, error) {
	for i := 0; i < len(ips); i++ {
		index := (l.currentId + i) % len(ips)
		if ips[index].IsActive {
			l.currentId = index + 1
			return ips[index].Value, nil
		}
	}
	return "", &NoActiveIpErr{}
}
