FROM alpine:3.10
#FROM nginx
VOLUME /opt
MAINTAINER smart.zhang

ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# wget http://tengine.taobao.org/download/tengine-2.2.3.tar.gz &&
# wget "https://ftp.pcre.org/pub/pcre/pcre-8.44.tar.gz"

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

RUN apk add wget gcc g++ make linux-headers



RUN mkdir /usr/local/src -p
#ADD tengine-2.2.3.tar.gz /usr/local/src/
ADD nginx-1.18.0.tar.gz /usr/local/src/
ADD pcre-8.44.tar.gz /usr/local/src/
ADD openssl-1.1.0i.tar.gz /usr/local/src/
ADD zlib-1.2.11.tar.gz /usr/local/src/
ADD perl-5.30.1.tar.gz /usr/local/src/

RUN ls -ah /usr/local/src/
#RUN ls -ah /usr/local/src/tengine-2.2.3
##RUN ls -alh /usr/local/src/pcre-8.44
#RUN ls -ah /usr/local/src/openssl-1.1.1m
#RUN ls -ah /usr/local/src/zlib-1.2.11

RUN cd /usr/local/src/perl-5.30.1 && ./Configure -des && make && make install
#    -Dprefix=$HOME/localperl \
RUN perl -v

# && tar -zxf tengine-2.2.3.tar.gz
RUN cd /usr/local/src && pwd && ls -alh nginx-1.18.0 && cd nginx-1.18.0 && ./configure --prefix=/usr/local/nginx --with-pcre=/usr/local/src/pcre-8.44 --with-ipv6 \
    --with-http_ssl_module \
    --with-zlib=/usr/local/src/zlib-1.2.11 \
    --with-openssl=/usr/local/src/openssl-1.1.0i \
    && make && make install

# 环境变量
ENV PATH $PATH:/usr/local/nginx/sbin

## 开放端口
#EXPOSE 80 443
#
## 工作目录
#WORKDIR /usr/local/nginx
#
## 执行命令
#CMD [ "nginx","-g","daemon off;" ]
