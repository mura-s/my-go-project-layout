FROM alpine

COPY myapp_linux /usr/bin/myapp
CMD ["myapp"]