FROM scratch
MAINTAINER Rosco Pecoltran <https://github.com/roscopecoltran>

EXPOSE 9000
COPY ./bin/krakend-admin-linux /app/krakend

ENV PATH=/app:$PATH
# VOLUME ["/app", "/conf.d/krakend"]

ENTRYPOINT ["/app/krakend"]
CMD ["-d", "-c", "/conf.d/krakend/default.yaml"]
