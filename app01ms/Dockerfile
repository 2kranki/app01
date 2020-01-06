
# vi:nu:et:sts=4 ts=4 sw=4

# Generate a minimal Docker container with app01ms installed

# This is a two step process. The first step is to generate the go
# binary program which will be passed to the second step. You have
# to keep a few things in mind when doing this. First, there are
# problems cross compiling go if cgo is needed and it is needed
# for some of the sql servers.  Second, you have to keep the glibc
# or its substitute at the same version level between the steps.
# Since Debian Linux and Ubuntu Linux are compatible, you just
# need to find the releases of each that have the same version
# of glibc. The Debian container is much larger than the Ubuntu
# one. So, Ubuntu was used.

# Alpine Linux was initially used for both steps. However, it is
# such a reduced Linux that there were problems compiling and
# running the code. When more time is available, this might be
# investigated further, because it is definitely much smaller
# than Ubuntu. But, it uses a glibc substitute.


# WARNING: This file must be located in the main directory
#           because of docker limitations.


# Build the application program using Debian-Golang container.
FROM golang:1.12-stretch AS golang
LABEL maintainer="bob@2kranki.us"
ENV GOLANG_DOCKER_CONTAINER=1

ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go get -u github.com/2kranki/jsonpreprocess \
    && go get -u github.com/2kranki/go_util \
    && go get -u github.com/shopspring/decimal \
    && go get -u "github.com/denisenkom/go-mssqldb"
WORKDIR /go/src/app01ms/cmd
COPY ./cmd      ./
WORKDIR /go/src/app01ms/pkg
COPY ./pkg      ./
WORKDIR /go/src/app01ms/vendor
COPY ./vendor   ./
WORKDIR /go/src/app01ms
COPY go.mod     ./
COPY go.sum     ./
RUN go build -o /go/bin/app01ms /go/src/app01ms/cmd/App01ms/*.go


# Create the Production Container with the program built in the
# prior step which is found in /usr/local/app of the golang con-
# tainer.
FROM ubuntu:cosmic AS production
# Note: we must keep glibc version the same as in the golang container.
LABEL maintainer="bob@2kranki.us"

RUN apt-get update && apt-get upgrade --yes
WORKDIR /usr/local/app
COPY --from=golang /go/bin/app01ms .
# Warning: COPY does not copy the actual directory, just its contents.
#           So, we must create those directories first, then COPY.
WORKDIR /usr/local/app/static
COPY ./static   .
WORKDIR /usr/local/app/tmpl
COPY ./tmpl     .
WORKDIR /usr/local/app
# The web server ip/port are different when running in the container
# vs batch.
ENV APP01MS_HTTP_SERVER="0.0.0.0"
ENV APP01MS_HTTP_PORT="8091"
EXPOSE 8091

CMD ["/usr/local/app/app01ms"]

