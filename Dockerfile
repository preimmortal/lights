FROM centos:7
COPY main /usr/local/bin/main
CMD [ "/usr/local/bin/main"]