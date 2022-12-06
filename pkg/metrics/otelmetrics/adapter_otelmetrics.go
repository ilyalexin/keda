/*
Copyright 2022 The KEDA Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package otelmetrics

import (
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
)

type OtelExporter struct{}

func New(endpoint string) {
	otlpOptions := []otlpmetrichttp.Option{}
	otlpOptions = append(otlpOptions, otlpmetrichttp.WithEndpoint(endpoint))
}

// func New(ctx context.Context, endpoint string, insecure bool, interval time.Duration) (*OtlpExporter, error) {
// 	otlpOptions := []http-exporter.Option{}
// 	otlpOptions = append(otlpOptions, http-exporter.WithEndpoint(endpoint))
// 	if insecure {
// 		otlpOptions = append(otlpOptions, http-exporter.WithInsecure())
// 	}
// 	// exp, err := otlpmetrichttp.New(ctx, otlpOptions...)
// 	// if err != nil {
// 	// 	log.Fatalf("Unable to initialize otlp http metrics exporter: %w", err)
// 	// 	return nil, err
// 	// }

// 	// readPeriod := sdkmetric.WithInterval(interval)

// 	// meterProvider := sdkmetric.NewMeterProvider(sdkmetric.WithReader(
// 	// 	sdkmetric.NewPeriodicReader(exp, readPeriod)))
// 	// //	defer func() {
// 	// // 	if err := meterProvider.Shutdown(ctx); err != nil {
// 	// // 		panic(err)
// 	// // 	}
// 	// // }()

// 	// global.SetMeterProvider(meterProvider)

// 	return &OtlpExporter{
// 		meterProvider: meterProvider,
// 	}, nil
// }
