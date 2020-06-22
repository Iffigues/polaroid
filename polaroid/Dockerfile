from golang:latest


WORKDIR /go/src/app
COPY ./polaroid .
RUN go mod init polaroid
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["go", "run", "."]
