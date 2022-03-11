# for nginx-bank 
FROM tengine-ipv6:0.1

MAINTAINER smart.zhang
#ENV RUN_USER nginx
#ENV RUN_GROUP nginx
ENV DATA_DIR /usr/share/nginx/html 
ENV LOG_DIR /var/log/nginx 

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

RUN mkdir /var/log/nginx -p 
RUN mkdir /data/web_files -p

RUN chown nginx.nginx -R /var/log/nginx 
 
ADD keys/7380947_nas.okzhang.com.key /usr/local/nginx/conf/7380947_nas.okzhang.com.key
ADD keys/7380947_nas.okzhang.com.pem /usr/local/nginx/conf/7380947_nas.okzhang.com.pem

ADD conf/nginx.conf /usr/local/nginx/conf/nginx.conf
ADD conf/web-api.conf /usr/local/nginx/conf/conf.d/web-api.conf
ADD conf/nginx.html /usr/share/nginx/html/
ADD conf/status.html /usr/local/nginx/conf/
#RUN mv /usr/local/nginx/conf/conf.d/default.conf /usr/local/nginx/conf/conf.d/default.conf_bak

RUN ls -alh /usr/local/nginx/conf/conf.d
RUN ls -ah /usr/local/nginx/

#EXPOSE 80
#EXPOSE 443
#ENTRYPOINT nginx -g "daemon off;"