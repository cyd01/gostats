# gostat

## Informations

**gostats** is a real-time Linux monitoring.

![](gostats.png)

The tool is based on famous [gopsutil front shirou](https://github.com/shirou/gopsutil).

## Building...

### gostats binary

```
$ go get
$ CGO_ENABLED=0 go build -a -installsuffix cgo
```

### Docker image

```
$ docker build . -t cyd01/gostats -f Dockerfile
```


## Running...

### The old way

```
$ PORT=80 ./startup.sh
Start gostats...
2020/06/22 19:09:14 Starting webserver on port 80 to directory /home/public/wwwroot
```

### The easy way

The easy is the the Docker way:

```
$ docker run --rm -t --privileged -p 80:80 \
	-e PORT=80 \
	-v /proc:/data/proc:ro \
	-v /sys:/data/sys:ro \
	-v /etc:/data/etc:ro \
	-v /var:/data/var:ro \
	-v /run:/data/run:ro \
	-v /dev:/data/dev:ro \
	cyd01/gostats
```

### The easiest

Much more easy, a compose file:

```
$ docker-compose up -d
Creating gostats_gostats_1 ... done
```
