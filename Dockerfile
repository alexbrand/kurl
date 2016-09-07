FROM golang:1.6
RUN curl https://glide.sh/get | sh
RUN go get github.com/alexbrand/kurl
WORKDIR $GOPATH/src/github.com/alexbrand/kurl
RUN glide up
RUN go build -o kurl main.go
RUN GOOS=windows GOARCH=amd64 go build -o kurl.exe
RUN chmod +x kurl
CMD ["./kurl"]
