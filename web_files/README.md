docker-compose -f dc-web-files.yml up


# TODO tengine with ipv6
cd /usr/local/src && wget http://tengine.taobao.org/download/tengine-2.2.3.tar.gz && tar -zxvf tengine-2.2.3.tar.gz && cd tengine-2.2.3 && ./configure --prefix=/usr/local/nginx --with-ipv6 && make && make install
