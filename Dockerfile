FROM golang:stretch

RUN mkdir -p /go/src/github.com/igormp/quiet-backend

WORKDIR /go/src/github.com/igormp/quiet-backend

RUN go get github.com/gorilla/mux

RUN go get gopkg.in/zabawaba99/firego.v1

COPY ./ /go/src/github.com/igormp/quiet-backend

RUN go build

CMD ["./quiet-backend"]