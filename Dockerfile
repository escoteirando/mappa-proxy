FROM golang:1.16
WORKDIR /go/src/github.com/guionardo/mappa_proxy/
RUN go get -u github.com/gin-gonic/gin
ARG GIN_MODE
COPY main.go .
COPY mappa/*.go ./mappa/
COPY go.* ./
RUN CGO_ENABLED=0 GOOS=linux go build -o mappa_proxy .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/guionardo/mappa_proxy/mappa_proxy .
CMD ["./mappa_proxy"]

