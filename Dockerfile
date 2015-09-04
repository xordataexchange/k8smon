FROM golang:1.5
MAINTAINER Brian Ketelsen <brian@xor.exchange>
RUN go get github.com/constabulary/gb/...
ADD . /app
WORKDIR /app
RUN gb build k8smon
CMD /app/bin/k8smon
