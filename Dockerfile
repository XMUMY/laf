FROM alpine

RUN apk add --no-cache ca-certificates

ADD bin/server /bin/server
ADD configs /configs/

EXPOSE 8000
EXPOSE 9000

ENTRYPOINT [ "/bin/server", "-conf", "/configs/config.yaml"]
