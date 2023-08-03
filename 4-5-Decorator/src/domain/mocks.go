package domain

import (
	"github.com/stretchr/testify/mock"
)

type NewMockHttpClient struct {
	mock.Mock
}

func (m NewMockHttpClient) SendRequest(req *HttpRequest) (*HttpResponse, error) {
	args := m.Called(req)
	return args.Get(0).(*HttpResponse), args.Error(1)
}
