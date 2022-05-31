#! /usr/bin/env bash

while true
do
    fswatch --one-event *.go static/*.css tmpl/*.html data/*.sql data/*.db > /dev/null
	pkill -9 cook
	go build -o cook cook.go
	./cook &
	#1 http :8081/connect
done 