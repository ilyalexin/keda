package otelmetrics

type MetricsProvider interface {
	RecordScalerMetric(namespace string, scaledObject string, scaler string, scalerIndex int, metric string, value float64)
	RecordScalerLatency(namespace string, scaledObject string, scaler string, scalerIndex int, metric string, value float64)
	RecordScalerActive(namespace string, scaledObject string, scaler string, scalerIndex int, metric string, active bool)
	RecordScalerError(namespace string, scaledObject string, scaler string, scalerIndex int, metric string, err error)
	RecordScaledObjectError(namespace string, scaledObject string, err error)

	IncrementTriggerTotal(triggerType string)
	DecrementTriggerTotal(triggerType string)
	IncrementCRDTotal(crdType, namespace string)
	DecrementCRDTotal(crdType, namespace string)
}
