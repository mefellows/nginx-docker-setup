FROM ubuntu:14.04

RUN apt-get update
RUN apt-get install -y libxml2 libxml2-dev libxml2-utils libaprutil1 libaprutil1-dev apache2-dev
RUN apt-get install -y wget tar
RUN mkdir /modsecurity
RUN mkdir /nginx

RUN wget -O /modsecurity/modsecurity-2.9.0.tar.gz https://www.modsecurity.org/tarball/2.9.0/modsecurity-2.9.0.tar.gz
RUN wget -O /nginx/ngx_openresty-1.7.7.2.tar.gz http://openresty.org/download/ngx_openresty-1.7.7.2.tar.gz

RUN cd /modsecurity && tar xvfz modsecurity-2.9.0.tar.gz
RUN cd /modsecurity/modsecurity-2.9.0 && ./configure --enable-standalone-module --disable-mlogc && make

RUN cd /nginx && tar -xvpzf ngx_openresty-1.7.7.2.tar.gz
RUN cd /nginx/ngx_openresty-1.7.7.2 && ./configure --add-module=/modsecurity/modsecurity-2.9.0/nginx/modsecurity \
                                                   --with-http_stub_status_module \
                                                   && make \
                                                   && make install

RUN cd /usr/local && ln -s openresty/nginx nginx
RUN echo "*       soft    nofile  3240000" >> /etc/security/limits.conf
RUN echo "*       hard    nofile  3240000" >> /etc/security/limits.conf

COPY nginx.conf /usr/local/nginx/conf/nginx.conf
COPY nginx-config.d/ /usr/local/nginx/conf/conf.d
COPY lib/uuid4.lua /usr/local/nginx/uuid4.lua

EXPOSE 80

VOLUME /var/log/nginx

# CMD /usr/local/nginx/sbin/nginx -g "daemon off;"
CMD /bin/echo 8096 > /writable-proc/sys/net/core/somaxconn && /usr/local/nginx/sbin/nginx -g "daemon off;"