docker run -d --name=aliyun_ddns_smart --restart=always --net=host \
    -e "AKID=" \
    -e "AKSCT=" \
    -e "DOMAIN=xxx.com" \
    -e "RR=nas" \
    -e "RECORDID=xxx" \
    -e "REDO=30" \
    -e "TTL=600" \
    -e "TIMEZONE=8.0" \
    -e "TYPE=AAAA" \
    registry.cn-shanghai.aliyuncs.com/jrdn/aliyun_ddns_app:1.0.0
