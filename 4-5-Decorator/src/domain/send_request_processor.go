package domain

type SendRequestProcessorI interface {
	SendRequest(req *HttpRequest) (*HttpResponse, error)
}

type SendRequestProcessor struct {
	next HttpClient
}
