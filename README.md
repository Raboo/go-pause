# Pause Container

This is a simple Go program that mimics the functionality of the `kubernetes/pause` container, similar to the one found in [Kubernetes](https://github.com/kubernetes/kubernetes/tree/master/build/pause).

## Description

The pause container is a minimalistic container that is used as a placeholder to keep a pod running on a node when no other containers are scheduled to run. This program acts as a placeholder by running indefinitely until it receives a termination signal.

## Usage

To run the pause container, simply build the Go program and execute the compiled binary. The program will start and print a log message indicating that the pause container has started. It will then wait for a termination signal. You can terminate the program by sending a SIGINT (Ctrl+C), SIGTERM, SIGQUIT (ofc SIGKILL as well) signal to it.

Compile in a local copy

```bash
go build -o pause main.go
./pause
```

Compile from Github in a Containerfile

```Dockerfile
FROM golang:1.22-alpine as build
RUN go install github.com/Raboo/go-pause@latest
FROM scratch
COPY --from=build /go/bin/go-pause /pause
USER 65535:65535
ENTRYPOINT ["/pause"]
```

## License

This project is licensed under the Unlicense - A license with no conditions whatsoever which dedicates works to the public domain. Unlicensed works, modifications, and larger works may be distributed under different terms and without source code.
See the LICENSE file for details.
