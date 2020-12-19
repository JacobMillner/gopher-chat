# Build the Image
FROM golang:alpine AS build-env
WORKDIR /go/src
ADD . /go/src/gopher-chat
ADD ../..
RUN cd /go/src/gopher-chat && go build -o main

# Run the Image
FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*
WORKDIR /go/src
COPY --from=build-env /go/src/gopher-chat /go/src
COPY ../../client /client
EXPOSE 8000
ENTRYPOINT [ “./gopher-chat” ]