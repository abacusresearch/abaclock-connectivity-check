FROM golang
ADD . /go/src/github.com/abacusresearch/abaclock-connectivity-check
RUN go get github.com/abacusresearch/abaclock-connectivity-check/...
RUN go install github.com/abacusresearch/abaclock-connectivity-check
ENTRYPOINT /go/bin/abaclock-connectivity-check
