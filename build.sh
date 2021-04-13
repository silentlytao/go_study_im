#!/bin/sh
rm -rf ./release
mkdir release
go build -o chat
chmod +x ./chat
cp chat ./release
cp favicon.ico ./release
cp -arf ./asset ./release
cp -arf ./view ./release

#注意一下 linux 运行 nohup ./chat >>./log.log 2>&1 &
