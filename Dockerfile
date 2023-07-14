FROM node:18-alpine AS build-ui

WORKDIR /build

COPY ./ui/v1/package.json .

RUN npm install

COPY ./ui/v1/ .

RUN npm run build


FROM golang:1.20-alpine AS build-server

WORKDIR /build

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN rm -rf ui/v1
COPY --from=build-ui /build/dist ./ui/v1/dist

RUN GOOS=linux go build -o ./ragbag ./cmd/server/main.go


FROM alpine:3.9 AS run

RUN apk update && apk add python3 g++ make bash git ffmpeg

WORKDIR /app

COPY --from=build-server /build/ragbag ./ragbag

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["/app/ragbag"]
