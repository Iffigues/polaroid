from golang:latest


WORKDIR /go/src/app
COPY ./staticserver/public .
RUN go mod init static
RUN go get -d -v ./...
RUN go install -v ./...
RUN rm .*
CMD ["go", "run", "."]
