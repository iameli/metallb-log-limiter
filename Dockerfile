FROM golang AS builder

WORKDIR /build
COPY ./log-limiter.go ./log-limiter.go
RUN go build ./log-limiter.go

FROM bitnami/metallb-speaker:0.9.6-debian-10-r26

COPY --from=builder /build/log-limiter /log-limiter
COPY ./entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
