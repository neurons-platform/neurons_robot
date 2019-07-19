#!/bin/bash


go build main.go
mv main dist/im_chatops_robot
cp -a ttf dist/
cp -a config dist/
cp -a start_robot.sh dist/
cp -a stop_robot.sh dist/

mkdir dist/tmp
rm dist.tar.gz
tar czvf dist.tar.gz dist

