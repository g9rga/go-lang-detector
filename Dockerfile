FROM golang:alpine3.11
ADD detector.go /opt
ADD go.mod /opt
WORKDIR /opt
RUN go build detector.go
CMD ["/opt/detector"]