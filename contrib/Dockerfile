FROM golang:1.21 AS builder

WORKDIR /app

COPY . ./

RUN make deps
RUN make build

FROM golang:1.21

WORKDIR /app

COPY static ./static
COPY templates ./templates
COPY config.yaml ./
COPY --from=builder /app/recall ./

EXPOSE 5555

CMD ["./recall"]
