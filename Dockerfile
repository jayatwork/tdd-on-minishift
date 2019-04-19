# My base image is pulling from Red Hat. We pass the Vendor supported criteria.
# This is just the Jenkins image because most images need a login on the redhat container catalog.
#FROM registry.access.redhat.com/devtools/go-toolset-1.10-rhel7

#EXPOSE 8080

# port src from tmp into image app-root
# These layers will only be re-built when Gopkg files are updated
#COPY /tmp/src /opt/app-root/src
#WORKDIR /opt/app-root/src

# non-root user
#USER 1001

#RUN go build -o main . 
#CMD ["./main"]

FROM golang:1.11

USER 1001

RUN mkdir -p /go/src/github.com/jayatwork/tdd-on-minishift
WORKDIR /go/src/github.com/jayatwork/tdd-on-minishift

COPY . /go/src/github.com/jayatwork/tdd-on-minishift
#RUN go get -u github.com/gorilla/mux
RUN go build

CMD ["./tdd-on-minishift"]

