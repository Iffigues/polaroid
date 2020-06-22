from golang:latest


WORKDIR /go/src/app
COPY ./staticserver .
RUN go mod init static
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["go", "run", "."]
