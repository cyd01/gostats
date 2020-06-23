#!/bin/sh

cd $(dirname $0)
export DIRNAME=$(pwd)
cd - > /dev/null
cd ${DIRNAME}/..
export DIRROOT=$(pwd)
cd - > /dev/null

export PATH=.:${PATH}
which gostats > /dev/null 2>&1 || {
	echo "Unable to find gostats !" >&2
	exit 1
}

DIR=${DIR:-/www/wwwroot}
if [ ! -d ${DIR} ] ; then
	if [ -d /www/wwwroot ] ; then DIR=/www/wwwroot ; fi
	if [ -d ${DIRNAME}/wwwroot ] ; then DIR=${DIRNAME}/wwwroot ; fi
	if [ -d $(pwd)/wwwroot ] ; then DIR=$(pwd)/wwwroot ; fi
	if [ -d /data/wwwroot ] ; then DIR=/data/wwwroot ; fi
fi

if [ ! -d ${DIR} ] ; then
	echo "Unable to find wwwroot directory" >&2
	exit 2
fi

PORT=${PORT:-80}

wpid=0
function clean() {
	if [ $wpid -gt 0 ] ; then echo "Killing gostats..." ; kill $wpid ; fi
}

if [ -d /data/proc ] ; then export HOST_PROC=/data/proc ; fi

if [ -d /data/sys ] ; then export HOST_SYS=/data/sys ; fi

if [ -d /data/etc ] ; then export HOST_ETC=/data/etc ; fi

if [ -d /data/var ] ; then export HOST_VAR=/data/var ; fi

if [ -d /data/run ] ; then export HOST_RUN=/data/run ; fi

if [ -d /data/dev ] ; then export HOST_DEV=/data/dev ; fi

trap clean 1 2 3 15

echo "Start gostats..."

gostats ${PORT} ${DIR} &
wpid=$!

if [ $# -ne 0 ] ; then
	"$@"
	clean
fi


wait
trap 1 2 3 15
