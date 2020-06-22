
.PHONY: all
all: compress image push clean

.PHONY: build
build: compile

compile: *.go
	go get
	CGO_ENABLED=0 go build -a -installsuffix cgo

.PHONY: compress
compress: build
	upx -9 gostats

.PHONY:	image
image: 
	docker build . -t cyd01/gostats -f Dockerfile

.PHONY: push
push:
	docker push cyd01/gostats

.PHONY: run
run:
	PORT=8585 ./startup.sh
	
.PHONY: docker
docker:
	docker run --rm -d --privileged --name gostats -p 8585:80 \
	-e PORT=80 \
	-v /proc:/data/proc:ro \
	-v /sys:/data/sys:ro \
	-v /etc:/data/etc:ro \
	-v /var:/data/var:ro \
	-v /run:/data/run:ro \
	-v /dev:/data/dev:ro \
	cyd01/gostats

.PHONY: stop
stop: 
	docker kill gostats

.PHONY: clean
clean:
	rm gostats
	docker kill gostats
	docker container prune -f
	docker image prune -f
	
