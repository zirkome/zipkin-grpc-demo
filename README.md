# Zipkin Demo with gRPC

> Tracing calls between microservices with Zipkin and gRPC

## Requirements

- Zipkin (https://github.com/openzipkin/zipkin)
- Optionnal: Zipkin Docker (https://github.com/openzipkin/docker-zipkin)
- govendor (https://github.com/kardianos/govendor)

## Installation

Locally with `govendor` (recommended):

```sh
govendor sync
go install ./cmd/...
```

With `go get`:

```sh
go get github.com/kokaz/zipkin-demo/cmd/...
```

## Run it!

The order matters (a bit) as some services depends on others.
Run the services in this order:

```sh
beta
centauri
alpha
```

And then use the client to talk to alpha:

```sh
alphaclient
```

Then go to your Zipkin UI (e.g. `localhost:9411`) to see the traces.
