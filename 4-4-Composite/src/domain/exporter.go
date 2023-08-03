package domain

type Exporter interface {
	export(message string) error
}
