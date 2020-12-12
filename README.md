# gostat

## Informations

**gostats** is a real-time Linux monitoring tool

![](gostats.png)

It can be used to monitor 
- cpu
- memory
- networks
- disks
- process

and very soon, Docker containers.

The tool is based on famous [gopsutil front shirou](https://github.com/shirou/gopsutil).

It is available on [Github](https://github.com/cyd01/gostats).

## Building...

### gostats binary

**gostats** developped in **Go**. In order to compile it after cloning the repository, just run

```
$ go get
$ CGO_ENABLED=0 go build -a -installsuffix cgo
```

### Docker image

**gpstats** is available in a **[Docker image](https://hub.docker.com/repository/docker/cyd01/gostats)**. To build the image:

```
$ docker build . -t cyd01/gostats -f Dockerfile
```

## Running...

### The "old" way

Simply run the process:

```
$ PORT=80 ./startup.sh
Start gostats...
2020/06/22 19:09:14 Starting webserver on port 80 to directory /home/public/wwwroot
```

### The easy way

The easy is the the Docker way:

```
$ docker run --rm -d -t --privileged -p 80:80 \
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
