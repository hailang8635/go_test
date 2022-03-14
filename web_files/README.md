
```
docker-compose -f dc-web-files.yml up
```

* centos方式 TODO

```
# TODO tengine with ipv6
cd /usr/local/src && wget http://tengine.taobao.org/download/tengine-2.2.3.tar.gz && tar -zxvf tengine-2.2.3.tar.gz && cd tengine-2.2.3 && ./configure --prefix=/usr/local/nginx --with-ipv6 && make && make install
```

* 基础镜像

```
# registry.cn-shanghai.aliyuncs.com/jrdn/alpine-ng:0.0.1
cd
docker build -it registry.cn-shanghai.aliyuncs.com/jrdn/alpine-ng:0.0.1 .
```