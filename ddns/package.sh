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

echo "start package"
ls -alh | grep aliyun
CGO_ENABLED=0 go build aliyun_ddns_app.go
echo "package done"
ls -alh | grep aliyun

docker build -t hailang8635/aliyun_ddns_app:$package_version .
docker push hailang8635/aliyun_ddns_app:$package_version

rm -rf aliyun_ddns_app