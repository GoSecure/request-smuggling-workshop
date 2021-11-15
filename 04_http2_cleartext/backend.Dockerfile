FROM golang:1.15-alpine

RUN apk add --no-cache git
WORKDIR /go/src/app
COPY src/*.go .
COPY src/*.htm .
COPY src/static/*.css ./static/
COPY src/static/*.png ./static/
COPY src/static/cover/*.jpg ./static/cover/

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
