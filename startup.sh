#!/bin/sh

cd $(dirname $0)
export DIRNAME=$(pwd)
cd - > /dev/null
cd ${DIRNAME}/..
export DIRROOT=$(pwd)
cd - > /dev/null

if [ -f ${DIRROOT}/../bin/gostats ] ; then
	export PATH=$PATH:.:${DIRROOT}/../bin
fi

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
