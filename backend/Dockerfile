# ---- Build Stage ----
FROM golang:1.20-alpine AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/server ./cmd/server

# ---- Run Stage ----
FROM alpine:3.18

EXPOSE 8080
COPY --from=build /go/bin/server /usr/local/bin/server
CMD ["server"]
