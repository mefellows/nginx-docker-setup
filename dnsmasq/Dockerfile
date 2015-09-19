FROM gliderlabs/alpine:3.1
RUN apk-install dnsmasq

COPY dnsmasq.conf /etc/

EXPOSE 5353

ENTRYPOINT ["dnsmasq", "-k"]