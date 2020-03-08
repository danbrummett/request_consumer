#! /bin/sh
BINNAME="request_consumer"
BUILDCONTAINER="build_request_consumer"
APPCONTAINER="dbrummett/request_consumer"
APPVERSION="latest"
docker build -f build_Dockerfile "$PWD" -t $BUILDCONTAINER
docker run --rm -v "$PWD":/$BINNAME $BUILDCONTAINER cp /app/$BINNAME /$BINNAME
docker build "$PWD" -t $APPCONTAINER:$APPVERSION
