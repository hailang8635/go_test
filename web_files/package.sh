#!/bin/bash
#eg:
# ./package_test.sh latest
# ./package_test.sh
# sudo docker login registry.cn-shanghai.aliyuncs.com
package_version=""

if [ -n "$1" ]
then
  package_version=$1
else
  package_version="0.0.1"
fi

echo "package $package_version .."
docker build -t registry.cn-shanghai.aliyuncs.com/jrdn/nginx-web-files:0.0.1 .
docker push registry.cn-shanghai.aliyuncs.com/jrdn/nginx-web-files:0.0.1
