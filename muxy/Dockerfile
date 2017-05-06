FROM debian:jessie

RUN apt-get update && apt-get install -y wget unzip iptables iproute net-tools sudo
RUN mkdir -p /opt/muxy/bin

WORKDIR /opt/muxy

# RUN wget https://github.com/mefellows/muxy/releases/download/v0.0.3/linux_amd64.zip?20170505 -O muxy.zip
RUN wget https://github.com/mefellows/muxy/releases/download/v0.0.4/linux_amd64.zip?20170505 -O muxy.zip
RUN unzip muxy.zip
RUN mv muxy /opt/muxy/bin/
RUN rm *.zip

ENV PATH /opt/muxy/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin

VOLUME ["/opt/muxy/conf"]

CMD ["muxy", "proxy", "--config", "/opt/muxy/conf/config.yml"]