FROM golang:latest AS builder

COPY . /build
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o service .

FROM scratch
COPY --from=builder /build/service ./
ENTRYPOINT ["./service"]
EXPOSE 8080