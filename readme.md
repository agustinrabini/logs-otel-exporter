## How to run

First get your credentials from your `Otel Endpoint`, `Instance ID` and `Password/API Token` from your Grafana Cloud account and replace the values on the `otel-collector.yml`.

Run the docker `docker build -t otel-collector . && docker run -p 4317:4317 -p 4318:4318 otel-collector`.

Execute the Go file. 

Validate on your Grafana Cloud instance that the Logs have been pushed.

## Errors

If something goes wrong the docker container will drop an error.