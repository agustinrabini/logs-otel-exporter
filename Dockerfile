FROM otel/opentelemetry-collector-contrib:latest

COPY otel-collector.yml /otel-config.yml

CMD ["--config", "/otel-config.yml"]