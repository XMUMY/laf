FROM alpine

RUN apk add --no-cache ca-certificates

ADD lost-found-service /lost-found-service

EXPOSE 9000

ENTRYPOINT [ "/lost-found-service" ]
