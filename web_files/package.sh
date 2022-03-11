#!/bin/bash
#eg:
# ./package_test.sh latest
# ./package_test.sh
# sudo docker login registry.cn-shanghai.aliyuncs.com

docker build -t registry.cn-shanghai.aliyuncs.com/jrdn/nginx-web-files:0.0.2 .
docker push registry.cn-shanghai.aliyuncs.com/jrdn/nginx-web-files:0.0.2
