#!/bin/bash
# 进程消失后重启
while true
do

if $(ps -ef |grep im_chatops_robot |grep -v grep 2>&1 >/dev/null)
then
   :
else
  nohup ./im_chatops_robot -f config/conf.ini 2>&1
fi

sleep 1

done

