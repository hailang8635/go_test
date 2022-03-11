FROM centos:7
MAINTAINER smart

# 添加文件
#COPY nginx.conf *.tar.gz /usr/src/

#RUN mkdir /usr/local/src -p
#ADD tengine-2.2.3.tar.gz /usr/local/src/
#ADD pcre-8.44.tar.gz /usr/local/src/
#ADD openssl-1.1.1m.tar.gz /usr/local/src/
#ADD zlib-1.2.11.tar.gz /usr/local/src/
##ADD perl-5.30.1.tar.gz /usr/local/src/

# 创建用户
RUN useradd -M -s /sbin/nologin nginx && \

# 安装依赖包
yum install wget deltarpm -y && \
wget -O /etc/yum.repos.d/epel-7.repo http://mirrors.aliyun.com/repo/epel-7.repo && \
yum -y install unzip bzip2 zlib zlib-devel pcre pcre-devel openssl openssl-devel \
geoip geoip-devel gd gd-devel gcc gcc-c++ make libtool