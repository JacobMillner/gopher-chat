# Container to build the binary
FROM golang:alpine AS build-env
WORKDIR /go/src
ADD . /go/src/gopher-chat
ADD ../..
RUN cd /go/src/gopher-chat && go build -o main

# Container to run the binary
FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*
WORKDIR /go/src
COPY --from=build-env /go/src/gopher-chat .
COPY ../../client /client
EXPOSE 8000
ENTRYPOINT [ “./gopher-chat” ]