FROM golang:latest as builder
ENV CGO_ENABLED="0" \
  GOOS="linux" \
  GOARCH="amd64"
WORKDIR /go/src/work
COPY . .
RUN make build

FROM alpine:latest
COPY --from=builder /go/src/work/app /app
ENTRYPOINT /app
