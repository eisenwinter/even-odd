FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM scratch
COPY --from=builder /build/main /app/
ENV EVEN_ODD_HTTP_PORT=8080
ENV EVEN_ODD_GRPC_PORT=8081
WORKDIR /app
CMD ["./main"]