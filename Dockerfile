FROM golang:alpine as builder

COPY . /root
RUN go build -o /root/signal /root/signal.go

FROM golang:alpine

COPY --from=builder /root/signal /

STOPSIGNAL SIGUSR2

EXPOSE 8080
CMD ["/signal"]