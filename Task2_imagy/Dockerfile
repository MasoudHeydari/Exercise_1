FROM golang:1.20-alpine AS base
WORKDIR /app

# Copy and build the app
COPY . .
RUN go build -o imagy main.go \
    && ls -l


FROM alpine:3.17
WORKDIR /imagy

COPY --from=base /app/imagy /imagy/
ENTRYPOINT ["./imagy", "run"]