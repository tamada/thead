FROM alpine:latest

ARG version=1.0.0

LABEL maintainer="Ryotaro Tanaka" \
      description="Notify me later!"

RUN    adduser -D thead \
    && apk --no-cache add --update --virtual .builddeps curl tar \
#    && curl -s -L -O https://github.com/tamada/nml/realeases/download/v${version}/nml-${version}_linux_amd64.tar.gz \
    # && curl -s -L -o nml-${version}_linux_amd64.tar.gz https://www.dropbox.com/s/3akz13eepvazu4m/nml-1.0.0_linux_amd64.tar.gz?dl=0 \
    && curl -s -L -o thead-${version}_linux_amd64.tar.gz https://www.dropbox.com/s/kojcoyh27vwu3uw/thead-1.0.0_linux_amd64.tar.gz?dl=0\
    && tar xfz thead-${version}_linux_amd64.tar.gz        \
    && mv thead-${version} /opt                           \
    && ln -s /opt/thead-${version} /opt/thead               \
    && ln -s /opt/thead-${version}/thead /usr/local/bin/thead \
    && rm thead-${version}_linux_amd64.tar.gz             \
    && apk del --purge .builddeps

ENV HOME="/home/thead"

WORKDIR /home/thead
USER    thead

ENTRYPOINT [ "/opt/thead/thead" ]