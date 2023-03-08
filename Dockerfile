FROM golang:1.20.1-alpine as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main main.go

FROM alpine
COPY --from=build /app/main .

ENTRYPOINT ["./main"]