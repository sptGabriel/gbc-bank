FROM golang:1.16.4

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN apt-get update && \
    apt-get install build-essential -y

CMD ["tail", "-f", "/dev/null"]

