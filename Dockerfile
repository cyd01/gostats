FROM	alpine

USER	root

RUN	mkdir /data
RUN	mkdir /www

COPY	gostats /bin/gostats
RUN	chmod +x /bin/gostats

COPY	wwwroot /www/wwwroot

COPY	Dockerfile /etc/Dockerfile

WORKDIR	/www

COPY	startup.sh /startup.sh
RUN	chmod +x /startup.sh

ENTRYPOINT [ "/startup.sh" ]

# cp -f /root/webserver . ; docker build . -t cyd01/serverinfo -f Dockerfile ; rm -f webserver
# docker run --rm -it --privileged -p 8585:80 -v /etc:/data/etc -v /proc:/data/proc -e PORT=80 cyd01/serverinfo 
