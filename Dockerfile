FROM golang:1.14.1-buster
USER root

WORKDIR /app

RUN apt-get update && apt-get upgrade -y
RUN apt-get install -y \
  vim less \
  zip unzip netcat

RUN go get -u \
  github.com/go-delve/delve/cmd/dlv \
  github.com/cespare/reflex

CMD ["/bin/bash"]
