FROM java:8

# create directory for gatling install
RUN mkdir -p /tmp/downloads
WORKDIR /tmp/downloads
RUN wget https://repo1.maven.org/maven2/io/gatling/highcharts/gatling-charts-highcharts-bundle/2.1.7/gatling-charts-highcharts-bundle-2.1.7-bundle.zip -O gatling.zip

RUN unzip gatling.zip -d /opt/
RUN mv /opt/gatling-charts-highcharts-bundle-2.1.7 /opt/gatling
RUN ls -la /opt/gatling
RUN rm *.zip
# change context to gatling directory
WORKDIR /opt/gatling

VOLUME ["/opt/gatling/conf","/opt/gatling/results","/opt/gatling/user-files"]

RUN ls /opt/gatling/

RUN echo "*       soft    nofile  65535" >> /etc/security/limits.conf
RUN echo "*       hard    nofile  65535" >> /etc/security/limits.conf

# Set environment variables
ENV PATH /opt/gatling/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
ENV GATLING_HOME /opt/gatling

ENTRYPOINT ["gatling.sh"]