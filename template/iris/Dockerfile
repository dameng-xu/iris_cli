FROM alpine:3.15 as prod

USER root
ENV TZ=Asia/Shanghai
ENV LANG=en_US.UTF-8
ENV LANGUAGE=en_US.UTF-8

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk update -vU --allow-untrusted && \
    apk add --no-cache tzdata

WORKDIR /home/moresec/{{.ModPrefix}}{{.Name}}

COPY  /release/bin/{{.ModPrefix}}{{.Name}} .

CMD ["./{{.ModPrefix}}{{.Name}}"]
