package observability

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
)

// InitMeter configures the OpenTelemetry meter provider with Prometheus exporter
func InitMeter(ctx context.Context) (*metric.MeterProvider, error) {
	exporter, err := prometheus.New()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize prometheus exporter: %w", err)
	}

	provider := metric.NewMeterProvider(
		metric.WithReader(exporter),
	)

	return provider, nil
}
