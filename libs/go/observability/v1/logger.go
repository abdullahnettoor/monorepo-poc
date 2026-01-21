package observability

import (
	"context"
	"log/slog"
	"os"

	"go.opentelemetry.io/otel/trace"
)

// InitLogger configures the standard slog logger with a handler that extracts trace IDs
func InitLogger() *slog.Logger {
	// Basic JSON handler
	handler := slog.NewJSONHandler(os.Stdout, nil)
	return slog.New(handler)
}

// LogWithTrace adds trace_id and span_id to the log record from the context
func LogWithTrace(ctx context.Context, logger *slog.Logger) *slog.Logger {
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return logger
	}

	return logger.With(
		slog.String("trace_id", span.SpanContext().TraceID().String()),
		slog.String("span_id", span.SpanContext().SpanID().String()),
	)
}
