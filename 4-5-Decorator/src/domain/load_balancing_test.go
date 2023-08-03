package domain_test

import (
	"4-5/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadBalancingSendRequest(t *testing.T) {
	mock := new(domain.NewMockHttpClient)

	expectedTargetIps := []string{"35.0.0.1", "35.0.0.2", "35.0.0.5", "35.0.0.1", "35.0.0.2"}
	client := domain.NewLoadBalancing(mock)
	ips := []domain.IP{
		{"35.0.0.1", true},
		{"35.0.0.2", true},
		{"35.0.0.3", false},
		{"35.0.0.4", false},
		{"35.0.0.5", true},
	}

	for i := 0; i < 5; i++ {
		request := domain.NewHttpRequest("http://test.com/abc")
		request.IPs = ips

		mock.On("SendRequest", request).Return(&domain.HttpResponse{Status: domain.Success}, nil).Once()
		_, _ = client.SendRequest(request)

		assert.Equal(t, expectedTargetIps[i], request.TargetIp)
	}
}
