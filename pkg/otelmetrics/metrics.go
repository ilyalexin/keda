package otelmetrics

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/instrument"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var log = logf.Log.WithName("otel_metrics_client")

const (
	DefaultMetricsNamespace = "keda"
)

type Provider struct {
	meter   metric.Meter
	metrics metrics
}

type metrics struct {
	scalerErrorsTotal syncint64.Counter
	scalerErrors      syncint64.Counter
}

func NewProvider(ctx context.Context, url string, insecure bool) (Provider, error) {
	otlpOptions := []otlpmetrichttp.Option{}
	otlpOptions = append(otlpOptions, otlpmetrichttp.WithEndpoint(url))
	if insecure {
		otlpOptions = append(otlpOptions, otlpmetrichttp.WithInsecure())
	}
	exporter, err := otlpmetrichttp.New(ctx, otlpOptions...)
	if err != nil {
		log.Error(err, "Unable to initialize otlp http metrics exporter")
		return Provider{}, err
	}
	meterProvider := sdkmetric.NewMeterProvider(sdkmetric.WithReader(
		sdkmetric.NewPeriodicReader(exporter)))

	meter := meterProvider.Meter("keda")
	metrics, err := initMetrics(meter)
	if err != nil {
		log.Error(err, "Unable to initialize otlp metrics")
	}

	return Provider{
		meter:   meter,
		metrics: metrics,
	}, nil
}

func initMetrics(meter metric.Meter) (metrics, error) {
	var m metrics
	scalerErrorsTotal, err := meter.SyncInt64().Counter(
		"errors_total",
		instrument.WithUnit("1"),
		instrument.WithDescription("Total number of errors for all scalers"),
	)
	if err != nil {
		log.Error(err, "Unable to setup scaler error totals metric")
		return m, err
	}
	m.scalerErrorsTotal = scalerErrorsTotal

	scalerErrors, err := meter.SyncInt64().Counter(
		"scaler_errors",
		instrument.WithUnit("1"),
		instrument.WithDescription("Number of scaler errors"),
	)
	if err != nil {
		log.Error(err, "Unable to setup scaler errors metric")
		return m, err
	}
	m.scalerErrors = scalerErrors
	return m, nil
}

func (p Provider) RecordScalerError(ctx context.Context, namespace string, scaledObject string, scaler string, scalerIndex int, metric string, err error) {
	if err != nil {
		attrs := setScalerAttribute(namespace, scaledObject, scaler, metric, scalerIndex)
		p.metrics.scalerErrors.Add(ctx, 1, attrs...)
		p.metrics.scalerErrorsTotal.Add(ctx, 1)
	}
}

func setScalerAttribute(namespace, scaledObject, scaler, metric string, scalerIndex int) []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.String("namespace", namespace),
		attribute.String("scaledObject", scaledObject),
		attribute.String("scaler", scaler),
		attribute.Int("scalerIndex", scalerIndex),
		attribute.String("metric", metric),
	}
}
