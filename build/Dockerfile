FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY cmd/main.go ./

RUN go mod download

RUN go build -o /go-image-compression

RUN mkdir -p /app/images

ENV IMAGES_FOLDER=/app/images 

ENTRYPOINT [ "/go-image-compression" ]
