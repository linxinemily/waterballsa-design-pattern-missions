package domain

import (
	"net/url"
)

type HttpRequest struct {
	Host     string
	Url      string
	TargetIp string
	IPs      []IP
}

func NewHttpRequest(urlString string) *HttpRequest {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return nil
	}

	host := parsedURL.Hostname()
	httpRequest := &HttpRequest{
		Host: host,
		Url:  urlString,
	}

	return httpRequest
}
