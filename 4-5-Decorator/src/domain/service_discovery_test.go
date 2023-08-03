package domain_test

import (
	"4-5/domain"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestServiceDiscoverySendRequest(t *testing.T) {
	tempFile, err := os.CreateTemp("", "test_service_discovery_config.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	content := []byte("test.com: 35.0.0.1, 35.0.0.2, 35.0.0.3, 35.0.0.4, 35.0.0.5")
	err = os.WriteFile(tempFile.Name(), content, 0644)
	if err != nil {
		t.Fatal(err)
	}

	mock := new(domain.NewMockHttpClient)

	responses := []struct {
		Response         *domain.HttpResponse
		Error            error
		ExpectedTargetIp string
	}{
		{&domain.HttpResponse{Status: domain.Success}, nil, "35.0.0.1"},
		{&domain.HttpResponse{Status: domain.Failure}, &domain.RequestFailedErr{}, "35.0.0.1"},
		{&domain.HttpResponse{Status: domain.Success}, nil, "35.0.0.2"},
		{&domain.HttpResponse{Status: domain.Failure}, &domain.RequestFailedErr{}, "35.0.0.2"},
		{&domain.HttpResponse{Status: domain.Failure}, &domain.RequestFailedErr{}, "35.0.0.3"},
	}

	client := domain.NewServiceDiscovery(tempFile.Name(), mock)

	for i := 0; i < 5; i++ {
		request := domain.NewHttpRequest("http://test.com/abc")

		mock.On("SendRequest", request).Return(responses[i].Response, responses[i].Error).Once()
		_, _ = client.SendRequest(request)

		assert.Equal(t, responses[i].ExpectedTargetIp, request.TargetIp)
	}
}
