package domain

type CompositeExporter struct {
	exporters []Exporter
}

func NewCompositeExporter(exporters []Exporter) *CompositeExporter {
	return &CompositeExporter{exporters}
}

func (c *CompositeExporter) export(message string) error {
	for _, exporter := range c.exporters {
		err := exporter.export(message)
		if err != nil {
			return err
		}
	}
	return nil
}
