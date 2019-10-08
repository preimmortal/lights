FROM centos:7
RUN yum install -y nmap iproute
COPY main /usr/local/bin/main
CMD ["/usr/local/bin/main"]