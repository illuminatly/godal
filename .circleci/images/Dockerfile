FROM ubuntu:19.04
MAINTAINER rfletcher@esri.com
RUN apt-get update && apt-get -y install git ssh tar gzip ca-certificates gdal-bin gdal-data libgdal-dev pkg-config
RUN wget https://dl.google.com/go/go1.11.5.linux-amd64.tar.gz && tar -xzvf go1.11.5.linux-amd64.tar.gz && mv go /usr/local
ENV GOROOT=/usr/local/go
ENV PATH=$GOROOT/bin:$PATH
ENTRYPOINT /bin/bash