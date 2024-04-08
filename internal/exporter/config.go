package exporter

import (
	"go.opentelemetry.io/collector/component"
)

type config struct {
}

var _ component.Config = (*config)(nil)

func (c *config) Validate() error {
	return nil

}
