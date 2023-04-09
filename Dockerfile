FROM golang:1.20-alpine as backend-dev

WORKDIR /app
RUN apk add build-base

# Copy go mod and sum files
COPY ./go.mod ./go.sum ./

# Download all required packages
RUN go mod download

COPY main.go .
COPY backend ./backend/
COPY docs ./docs/

RUN CGO_ENABLED=1 GOOS=linux go build -o mappa_proxy  -ldflags="-X 'main.Build=$(date +%Y-%m-%dT%H:%M:%S%z)'" .


FROM alpine:latest as final-build
# FROM cgr.dev/chainguard/static:latest as final-build
# Ensure updated CA certificates
RUN apk --no-cache add ca-certificates

# needed for timezones
RUN apk add --no-cache tzdata

WORKDIR /root/
COPY --from=backend-dev /app/mappa_proxy .
CMD ["./mappa_proxy"]
