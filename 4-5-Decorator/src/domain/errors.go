package domain

import (
	"fmt"
)

type InBlackListErr struct {
	ip string
}

func NewInBlackListErr(ip string) *InBlackListErr {
	return &InBlackListErr{ip: ip}
}

func (i *InBlackListErr) Error() string {
	return fmt.Sprintf("IP address %s is blacklisted", i.ip)
}

type RequestFailedErr struct{}

func (i *RequestFailedErr) Error() string {
	return "send request failed"
}

type NoActiveIpErr struct{}

func (i *NoActiveIpErr) Error() string {
	return "no active ip"
}
