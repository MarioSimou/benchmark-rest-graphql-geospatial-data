FROM golang:1.13-buster

ENV APP=app

WORKDIR /go/src/github.com/MarioSimou/${APP}
COPY . .
RUN go get github.com/cespare/reflex \
    && go get -v ./...

EXPOSE 3000
CMD reflex -s -r '\.go$' -d fancy -- go run ./cmd/gis/main.go