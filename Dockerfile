FROM centos:7
WORKDIR /root/go/src/github.com/preimmortal/smarthome
RUN mkdir -p /root/go/src/github.com/preimmortal/smarthome
RUN yum update -y
RUN yum install -y nmap iproute iproute2 git wget gcc
RUN wget -P /bin https://storage.googleapis.com/golang/go1.10.1.linux-arm64.tar.gz
RUN tar -xvzf /bin/go1.10.1.linux-arm64.tar.gz -C /bin
ENV GOROOT /bin/go
ENV GOPATH /root/go
ENV PATH /bin/go/bin:$PATH
COPY . /root/go/src/github.com/preimmortal/smarthome
RUN go get github.com/Ullaakut/nmap github.com/cenkalti/backoff github.com/gorilla/handlers github.com/gorilla/mux github.com/grandcat/zeroconf github.com/hashicorp/go-memdb github.com/miekg/dns github.com/pkg/errors github.com/stretchr/testify
RUN go build /root/go/src/github.com/preimmortal/smarthome/cmd/smarthome/main.go
CMD ["/root/go/src/github.com/preimmortal/smarthome/main"]
