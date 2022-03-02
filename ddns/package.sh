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
  package_version="1.0.0"
fi

echo "package $package_version .."

CGO_ENABLED=0 go build aliyun_ddns_app.go

docker build -t registry.cn-shanghai.aliyuncs.com/jrdn/aliyun_ddns_app:$package_version .
docker push registry.cn-shanghai.aliyuncs.com/jrdn/aliyun_ddns_app:$package_version

rm -rf aliyun_ddns_app