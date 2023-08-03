package domain

type HttpClient interface {
	SendRequest(req *HttpRequest) (response *HttpResponse, err error)
}
