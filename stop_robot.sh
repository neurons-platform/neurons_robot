#!/bin/bash

ps -ef |grep start_robot.sh |grep -v grep  |awk '{print $2}' |xargs -i kill -9 {}
ps -ef |grep im_chatops_robot |grep -v grep  |awk '{print $2}' |xargs -i kill -9 {}
