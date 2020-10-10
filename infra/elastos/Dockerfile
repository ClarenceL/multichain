FROM ubuntu:xenial

RUN apt-get update && apt-get install
RUN apt-get install --yes curl

# this creates a folder called elastos-ela-v0.6.0
RUN curl https://download.elastos.org/elastos-ela/elastos-ela-v0.6.0/elastos-ela-v0.6.0-linux-x86_64.tgz | tar -xz
RUN mv ./elastos-ela-v0.6.0/* /root

COPY config.json /root/
COPY keystore.dat /root/
COPY blockdata.tar.gz /root/
RUN tar -xzf /root/blockdata.tar.gz -C /root/

# set workdir so that config.json is properly imported
# don't run /root/ela, it can't find config.json
WORKDIR /root

EXPOSE 20333-20338

ENTRYPOINT ["/bin/sh", "-c", "./ela"]