#!/bin/sh

killall main && echo "Sent kill"
rm -f ./main && echo "Removed old binary"
echo "Going to build"
go build cmd/service/main.go
echo "Going to run..."
nohup ./main > ./out.log &
echo "Ran successfully"