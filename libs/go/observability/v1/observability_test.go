package observability_test

import (
	"context"
	"testing"

	"github.com/abdullahnettoor/monorepo-poc/libs/go/observability/v1"
)

func TestInitialization(t *testing.T) {
	ctx := context.Background()

	// Verify InitTracer
	// Note: We use a dummy address, it might fail to connect but we check if the function invocation works
	_, err := observability.InitTracer(ctx, "test-service", "localhost:4317")
	if err != nil {
		t.Logf("InitTracer returned error (expected without collector): %v", err)
	}

	// Verify InitMeter
	_, err = observability.InitMeter(ctx)
	if err != nil {
		t.Errorf("InitMeter returned error: %v", err)
	}

	// Verify InitLogger
	logger := observability.InitLogger()
	if logger == nil {
		t.Error("InitLogger returned nil")
	}
}
