package domain

type Status int

const (
	Success Status = iota
	Failure
)

type HttpResponse struct {
	Status Status
}
