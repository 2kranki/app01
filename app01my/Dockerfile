
# Generate a minimal Docker container with app01my installed
# vi:nu:et:sts=4 ts=4 sw=4

# Build the application program.
FROM golang:1.12-stretch AS golang
LABEL maintainer="bob@2kranki.us"
ENV GOLANG_DOCKER_CONTAINER=1

ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go get github.com/2kranki/jsonpreprocess \
    && go get github.com/2kranki/go_util \
    && go get github.com/shopspring/decimal \
    && go get "github.com/go-sql-driver/mysql"
WORKDIR /go/src
COPY ./src/* ./
RUN go build -o /go/bin/app01my .                                       

# We put the application program built in the prior step into
# /usr/local/app including all of its support files.
FROM ubuntu:cosmic AS production
# Note: we must keep glibc version the same as in the golang container.
LABEL maintainer="bob@2kranki.us"

RUN apt-get update && apt-get upgrade --yes
WORKDIR /usr/local/app
COPY --from=golang /go/bin/app01my .
# Warning: COPY does not copy the actual directory, just its contents.
#           So, we must create those directories first, then COPY.
WORKDIR /usr/local/app/css
COPY ./css      .
WORKDIR /usr/local/app/html
COPY ./html     .
WORKDIR /usr/local/app/static
COPY ./static   .
WORKDIR /usr/local/app/tmpl
COPY ./tmpl     .
WORKDIR /usr/local/app
# The web server ip/port are different when running in the container
# vs batch.
ENV APP01MY_HTTP_SERVER="0.0.0.0"
ENV APP01MY_HTTP_PORT="8092"
EXPOSE 8092

CMD ["/usr/local/app/app01my"]

