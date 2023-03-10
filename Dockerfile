FROM golang:1.20.1-alpine as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main main.go

FROM alpine
WORKDIR /app
COPY --from=build /app/main /app/main
COPY --from=build /app/config /app/config 
ENV ENV=prod
ENTRYPOINT ["/app/main"]