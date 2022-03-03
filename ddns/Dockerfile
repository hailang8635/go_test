FROM alpine:3.10
VOLUME /opt
MAINTAINER smart.zhang

ADD aliyun_ddns_app /usr/local/bin/aliyun_ddns_app
RUN chmod +x /usr/local/bin/aliyun_ddns_app

RUN mkdir /lib64
RUN ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENTRYPOINT aliyun_ddns_app ${AKID} ${AKSCT} ${RR} ${RECORDID} ${TYPE} ${TTL}
#CMD ["sleep", "3600"]

