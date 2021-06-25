FROM    golang as buildbinary

RUN     mkdir -p /go/src/gostats
COPY    . /go/src/gostats
RUN     cd /go/src/gostats && rm go.mod go.sum 2> /dev/null || test 1
RUN     cd /go/src/gostats && go mod init gostats && go get && GOOS=linux GOARCG=386 CGO_ENABLED=0 go build -a -installsuffix cgo && ./gostats


FROM    alpine as buildcompress
COPY    --from=buildbinary /go/src/gostats/gostats /usr/local/bin/gostats
RUN     apk add --no-cache upx=3.96-r1
RUN     cd /usr/local/bin && upx --best gostats


FROM    busybox as buildimage

COPY    --from=buildcompress /usr/local/bin/gostats /usr/local/bin/gostats
RUN     chmod +x /usr/local/bin/gostats
COPY	wwwroot /web/www
RUN     mkdir -p /web/www


FROM    scratch

COPY    --from=buildimage / /

EXPOSE  80

WORKDIR /web

HEALTHCHECK	--interval=30s --timeout=15s --start-period=15s --retries=3 CMD wget --spider http://localhost

CMD     [ "/usr/local/bin/gostats", "80", "/web/www" ]
