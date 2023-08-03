package domain

import (
	"math/rand"
	"time"
)

type FakeHttpClient struct {
}

func NewFakeHttpClient() *FakeHttpClient {
	return &FakeHttpClient{}
}

func (c FakeHttpClient) SendRequest(req *HttpRequest) (*HttpResponse, error) {
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(2)

	if result == 1 {
		return &HttpResponse{Status: Success}, nil
	}

	return &HttpResponse{Status: Failure}, &RequestFailedErr{}
}
