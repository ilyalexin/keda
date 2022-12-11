package otelmetrics

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel/metric/global"
)

func New(ctx context.Context, endpoint string, insecure bool, interval time.Duration) (*OtlpExporter, error) {
	otlpOptions := []otlpmetrichttp.Option{}
	otlpOptions = append(otlpOptions, otlpmetrichttp.WithEndpoint(endpoint))
	if insecure {
		otlpOptions = append(otlpOptions, otlpmetrichttp.WithInsecure())
	}
	exp, err := otlpmetrichttp.New(ctx, otlpOptions...)
	if err != nil {
		log.Fatalf("Unable to initialize otlp http metrics exporter: %w", err)
		return nil, err
	}

	readPeriod := sdkmetric.WithInterval(interval)

	meterProvider := sdkmetric.NewMeterProvider(sdkmetric.WithReader(
		sdkmetric.NewPeriodicReader(exp, readPeriod)))
	//	defer func() {
	// 	if err := meterProvider.Shutdown(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()

	global.SetMeterProvider(meterProvider)

	return &OtlpExporter{
		meterProvider: meterProvider,
	}, nil
}

func (exporter *OtlpExporter) initMetrics(meterName string) {
	meter = exporter.meterProvider.Meter(meterName)
	scalerErrorsTotal, _ := meter.SyncInt64().Counter(
		"errors_total",
		instrument.WithUnit("1"),
		instrument.WithDescription("Total number of errors for all scalers"),
	)

	scalerMetricsValue, _ := meter.AsyncFloat64().Gauge(
		"metrics_value",
		instrument.WithUnit("1"),
		instrument.WithDescription("Metric Value used for HPA"),
	)
	meter.RegisterCallback([]instrument.Asynchronous{scalerMetricsValue}, func(ctx context.Context) {

	})
}
