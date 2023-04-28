FROM golang:1.20 as builder
COPY ../../pkg   /go/src/ynab-metrics/pkg/
COPY ../../go.mod /go/src/ynab-metrics/
COPY ../../go.sum /go/src/ynab-metrics/
COPY ../../cmd/ynab-metrics/. /go/src/ynab-metrics/
WORKDIR /go/src/ynab-metrics/
RUN go build -o target/ynab-metrics -v

from ubuntu:22.04 as add-ca-certs
RUN apt-get update
RUN apt-get install ca-certificates -y

FROM ubuntu:22.04
COPY --from=builder /go/src/ynab-metrics/target/ynab-metrics /usr/local/bin/ynab-metrics
COPY --from=builder /go/src/ynab-metrics/target/ynab-metrics /
COPY --from=add-ca-certs /etc/ssl /etc/ssl
EXPOSE 8080
CMD ynab-metrics --config $TOKEN_FILE_PATH
