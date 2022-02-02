##
## Build
##

FROM golang:1.17-alpine as build

WORKDIR /app

COPY go.mod ./

COPY cmd/main.go ./

RUN go mod download

RUN go build -o /app/go-image-compression

##
## Deploy
##

FROM alpine:3.15

ENV IMAGES_FOLDER=/app/images 

WORKDIR /app

RUN mkdir -p /app/images

COPY --from=build /app/go-image-compression /app/go-image-compression

ENTRYPOINT [ "/app/go-image-compression" ]


