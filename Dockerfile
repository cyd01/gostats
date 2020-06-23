FROM	alpine

USER	root

RUN	mkdir /data
RUN	mkdir /www

COPY	gostats /bin/gostats
RUN	chmod +x /bin/gostats

COPY	wwwroot /www/wwwroot

COPY	Dockerfile /etc/Dockerfile

EXPOSE	80
WORKDIR	/www

COPY	startup.sh /startup.sh
RUN	chmod +x /startup.sh

ENTRYPOINT [ "/startup.sh" ]
