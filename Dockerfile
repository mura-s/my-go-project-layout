FROM alpine

COPY myapp /usr/local/bin/
ENTRYPOINT ["myapp"]