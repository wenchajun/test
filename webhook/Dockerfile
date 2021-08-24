# Build the manager binary
FROM golang:1.14 as builder

WORKDIR /
COPY main.go /go/src/main.go
COPY go.mod go.mod
COPY go.sum go.sum
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.io,direct
RUN  go mod download

RUN CGO_ENABLED=0 go build /go/src/main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /main .
USER nonroot:nonroot

ENTRYPOINT ["/main"]

