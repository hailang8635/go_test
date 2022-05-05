# for nginx-bank 
FROM nginx 
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
 
ADD keys/7380947_nas.okzhang.com.key /etc/nginx/7380947_nas.okzhang.com.key
ADD keys/7380947_nas.okzhang.com.pem /etc/nginx/7380947_nas.okzhang.com.pem

ADD conf/nginx.conf /etc/nginx/nginx.conf
ADD conf/web-api.conf /etc/nginx/conf.d/web-api.conf
ADD conf/nginx.html /usr/share/nginx/html/
ADD conf/status.html /etc/nginx/
RUN mv /etc/nginx/conf.d/default.conf /etc/nginx/conf.d/default.conf_bak

RUN ls -ah /etc/nginx
RUN ls -alh /etc/nginx/conf.d

EXPOSE 80 
EXPOSE 443 
ENTRYPOINT nginx -g "daemon off;"