FROM alpine:3.6
MAINTAINER Rosco Pecoltran <https://github.com/roscopecoltran>

# Install Gosu to /usr/local/bin/gosu
ARG GOSU_VERSION=${GOSU_VERSION:-"1.10"}
ADD https://github.com/tianon/gosu/releases/download/${GOSU_VERSION}/gosu-amd64 /usr/local/sbin/gosu

EXPOSE 8080
COPY ./bin/krakend-gin-linux /app/krakend

# Install runtime dependencies & create runtime user
RUN chmod +x /usr/local/sbin/gosu \
 && apk --no-cache --no-progress add ca-certificates \
 && adduser -D app -h /data -s /bin/sh

# Container configuration
ENV PATH=/app:$PATH
VOLUME ["/app", "/conf.d/krakend", "/data"]

# ENTRYPOINT ["/usr/local/sbin/gosu", "app", "/app/krakend"]
ENTRYPOINT ["/app/krakend"]
CMD ["-d", "-c", "/conf.d/krakend/default.yaml"]
