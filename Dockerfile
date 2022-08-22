FROM node:14-alpine as frontend-dev

WORKDIR /app

COPY frontend/package.json ./
RUN yarn install
RUN npx browserslist@latest --update-db
RUN yarn global add @quasar/cli

FROM frontend-dev as frontend-build

COPY frontend/. . 
RUN yarn
RUN quasar build


FROM golang:1.19-alpine as backend-dev

WORKDIR /app
# Copy go mod and sum files
COPY ./go.mod ./go.sum ./

# Download all required packages
RUN go mod download


# RUN go get -u github.com/gin-gonic/gin
# ARG GIN_MODE
COPY main.go .
COPY backend ./backend/
COPY docs ./docs/

RUN apk add build-base
# CGO_ENABLED=0
RUN CGO_ENABLED=1 GOOS=linux go build -o mappa_proxy . 

FROM alpine:latest as final-build
# Ensure updated CA certificates
RUN apk --no-cache add ca-certificates

# needed for timezones
RUN apk add --no-cache tzdata

WORKDIR /root/
COPY --from=backend-dev /app/mappa_proxy .
COPY --from=frontend-build /web ./web/
CMD ["./mappa_proxy"]
