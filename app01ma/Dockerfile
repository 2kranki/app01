
# Generate a minimal Docker container with app01ma installed            

FROM golang:latest AS golang
LABEL maintainer="bob@2kranki.us"
ENV GOLANG_DOCKER_CONTAINER=1
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN apt-get update && apt-get install --yes sqlite3 libsqlite3-dev 
RUN go get github.com/2kranki/jsonpreprocess \
    && go get github.com/2kranki/go_util \
    && go get github.com/shopspring/decimal \
    && go get "github.com/go-sql-driver/mysql"
WORKDIR /go/src
COPY ./src/* ./
RUN go build -o /go/bin/app01ma .                                       
RUN ls -al /go/bin/

FROM alpine:latest AS production
WORKDIR /usr/local/app/
COPY --from=golang /go/bin/app01ma .
COPY ./css      ./
COPY ./html     ./
COPY ./static   ./
COPY ./tmpl     ./
EXPOSE 8090
CMD ["/usr/local/app/app01ma"]

