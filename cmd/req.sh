#!/bin/sh

case "$1" in
	body)
	curl -X POST 127.0.0.1:8002/v1/body -d "<div id='lol'><img src='' href='asd'/><td></td><p> asd</p><div> lol</div></div>"
	;;
	title)
	curl -X POST 127.0.0.1:8002/v1/title -d "hello what is this ????? ||||| <<<<"
	;;
esac
