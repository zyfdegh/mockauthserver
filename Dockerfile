FROM alpine
MAINTAINER zyfdegh <zyfdegg@gmail.com>

# fix library dependencies
# otherwise golang binary may encounter 'not found' error
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

COPY mockauthserver /mockauthserver
CMD /mockauthserver
