# Build stage
FROM golang:1.21.3 AS build-env
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api-server/main.go

# Final stage
FROM alpine:latest
ARG ENV_FILE=.env
RUN addgroup -S app && adduser -S app -G app
WORKDIR /home/app
COPY --from=build-env /app/main .
COPY --from=build-env /app/${ENV_FILE} .env
RUN chown -R app:app /home/app
USER app
CMD ["./main"]