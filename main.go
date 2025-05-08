package main

import (
	"context"
	"fmt"
	logs "log"
	"time"

	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	otellog "go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/log/global"
	sdklog "go.opentelemetry.io/otel/sdk/log"
)

// docker build -t otel-collector . && docker run -p 4317:4317 -p 4318:4318 otel-collector
// docker stop $(docker ps -a -q)
// docker rm $(docker ps -a -q)
func main() {
	ctx := context.Background()

	// Configure OTLP log exporter
	exporter, err := otlploghttp.New(ctx, otlploghttp.WithEndpointURL("http://localhost:4318/v1/logs"))
	if err != nil {
		logs.Fatalf("Failed to initialize OTLP log exporter: %v", err)
	}

	// Create log provider
	processor := sdklog.NewBatchProcessor(exporter)
	provider := sdklog.NewLoggerProvider(sdklog.WithProcessor(processor))
	global.SetLoggerProvider(provider)

	// Emit a log record
	record := otellog.Record{}
	record.SetSeverity(otellog.SeverityInfo)                        // Set severity level
	record.SetTimestamp(time.Now())                                 // Set timestamp
	record.SetBody(otellog.StringValue("OTLP logging is working!")) // Log message
	record.AddAttributes(otellog.String("status", "successful"))
	// Create logger from provider
	logger := provider.Logger("otel-logger")

	// Emit log entry
	logger.Emit(ctx, record)

	err = provider.Shutdown(ctx)
	if err != nil {
		logs.Fatalf("Failed to shutdown logger provider: %v", err)
	}

	fmt.Println("Log sent successfully!")

}
