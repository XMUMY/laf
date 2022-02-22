FROM alpine

RUN apk add --no-cache ca-certificates

ADD lost_found /bin/lost_found
ADD configs /configs/

EXPOSE 8000
EXPOSE 9000

ENTRYPOINT [ "/bin/lost_found", "-conf", "/configs/config.yaml"]
