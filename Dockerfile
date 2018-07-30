FROM centos

RUN yum install -y git golang
ENV GOPATH /root/go
ENV REPO github.com/delphinus/gosampleserver
RUN git clone https://$REPO $GOPATH/src/$REPO
WORKDIR $GOPATH/src/$REPO
RUN git fetch && git checkout e0c4b41
RUN go get ./...
RUN go build
ENTRYPOINT ./gosampleserver
