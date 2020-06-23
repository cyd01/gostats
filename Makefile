COMPRESS=upx
ECHO=echo
EXE=gostats
IMAGE=cyd01/gostats

.PHONY: info
info:
	@$(ECHO) "Avalable commands"
	@$(ECHO) "- build		to compile the gostats binary"
	@$(ECHO) "- compress	to compress the gostats binary (upx needed)"
	@$(ECHO) "- image		to make the Docker image"
	@$(ECHO) "- push		to push the Docker image on the Docker hub"
	@$(ECHO) "- clean		to ... clean"
	@$(ECHO) "- all		build+compress+image+push"
	@$(ECHO) "- run		to simply run the gostats binary"
	@$(ECHO) "- docker		to run the Docker image in background container"
	@$(ECHO) "- stop		to stop a background container"

.PHONY: all
all: compress image push clean

.PHONY: build
build: compile

compile: $(EXE)

$(EXE): *.go
	go get
	CGO_ENABLED=0 go build -a -installsuffix cgo

.PHONY: compress
compress: $(EXE)
	$(COMPRESS) -9 $(EXE)

.PHONY:	image
image: 
	docker build . -t $(IMAGE) -f Dockerfile

.PHONY: push
push:
	docker push $(IMAGE)

.PHONY: run
run:
	PORT=8585 ./startup.sh
	
.PHONY: docker
docker:
	docker run --rm -d --privileged --name $(EXE) -p 8585:80 \
	-e PORT=80 \
	-v /proc:/data/proc:ro \
	-v /sys:/data/sys:ro \
	-v /etc:/data/etc:ro \
	-v /var:/data/var:ro \
	-v /run:/data/run:ro \
	-v /dev:/data/dev:ro \
	$(IMAGE)

.PHONY: stop
stop: 
	docker kill $(EXE)

.PHONY: clean
clean:
	rm $(EXE)
	docker kill $(EXE)
	docker container prune -f
	docker image prune -f
	
