package domain

import (
	"fmt"
	"os"
	"strings"
)

type Blacklist struct {
	next HttpClient
	ips  []string
}

func NewBlacklist(configFile string, next HttpClient) *Blacklist {
	ips, err := readIPAddressesFromFile(configFile)
	if err != nil {
		panic("parse blacklist config file error")
	}

	return &Blacklist{next, ips}
}

func (l *Blacklist) SendRequest(req *HttpRequest) (*HttpResponse, error) {
	if l.isInBlacklist(req.TargetIp) {
		fmt.Printf("[Blacklist] IP address %s is blacklisted\n", req.TargetIp)
		return nil, NewInBlackListErr(req.TargetIp)
	}

	return l.next.SendRequest(req)
}

func (l *Blacklist) isInBlacklist(ip string) bool {
	for _, blacklistedIP := range l.ips {
		if ip == blacklistedIP {
			return true
		}
	}
	return false
}

func readIPAddressesFromFile(configFile string) ([]string, error) {
	content, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	ipAddresses := strings.Split(string(content), ",")
	for i := 0; i < len(ipAddresses); i++ {
		ipAddresses[i] = strings.TrimSpace(ipAddresses[i])
	}

	return ipAddresses, nil
}
