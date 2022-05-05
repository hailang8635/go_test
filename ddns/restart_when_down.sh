#!/bin/bash
pcnt=`ps -ef |grep frp |grep -v grep | awk '{print $2}' |wc -l`

echo " pid count ---> "$pcnt

if [ "$pcnt" != 0 ];then
        echo "process is running "
else
        echo "process not running"
        #启动 process
        echo "开始启动 process...."
        sh start.sh
        sleep 3; 
        ps -ef|grep frp |grep -v grep
        echo "process 启动成功"
        tail -f frp.log
fi

echo $PID
