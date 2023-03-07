FROM golang:1.20.1 as build
WORKDIR /app
COPY . .
RUN go build -o app main.go

# FROM alpine
# COPY --from=build /go-app/app   /app

ENTRYPOINT ["./app"]