FROM alpine:3.4

ENV GOPATH /go
ENV PATH $PATH:/go/bin

RUN echo "@community http://nl.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories

RUN apk update && apk upgrade && \
    mkdir -p /go/bin && \
    apk -v add git make bash go@community musl-dev curl && \
    go version

COPY ./ /go/src/github.com/appcelerator/ampbeat

RUN cd $GOPATH/src/github.com/appcelerator/ampbeat && \
    make && \
    echo ampbeat built && \
    mkdir -p /etc/ampbeat && \
    cp $GOPATH/src/github.com/appcelerator/ampbeat/ampbeat /etc/ampbeat && \
    cp $GOPATH/src/github.com/appcelerator/ampbeat/ampbeat-confimage.yml /etc/ampbeat/ampbeat.yml && \
    cp $GOPATH/src/github.com/appcelerator/ampbeat/*.json /etc/ampbeat && \
    chmod +x /etc/ampbeat/ampbeat && \
    cd $GOPATH && \
    rm -rf $GOPATH/src && \
    rm -rf /root/.glide

WORKDIR /etc/ampbeat

CMD ["/etc/ampbeat/ampbeat", "-e"]
