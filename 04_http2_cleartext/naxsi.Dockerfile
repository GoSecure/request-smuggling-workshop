FROM ubuntu/nginx:1.18-20.04_edge


RUN mkdir /setup-naxsi
WORKDIR /setup-naxsi

ENV NAXSI_PACKAGE=ubuntu-focal-libnginx-mod-http-naxsi_1.3_amd64.deb
COPY deb/${NAXSI_PACKAGE} /setup-naxsi/${NAXSI_PACKAGE}

RUN apt-get install -f /setup-naxsi/${NAXSI_PACKAGE}

# ENTRYPOINT /bin/sh