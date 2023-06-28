FROM harbor.moresec.cn/sast/golang:1.18.2-alpine3.15-amd64 as builder

COPY . /workdir
WORKDIR /workdir/{{.ModPrefix}}{{.Name}}

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-w -s" -o {{.ModPrefix}}{{.Name}} ./cmd/main.go

FROM harbor.moresec.cn/moresec/alpine_cn:1.0.0

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
  apk update -vU --allow-untrusted && \
  apk add --no-cache bash curl iputils tcptraceroute iproute2 vim python2 p7zip openssl

ENV ZONEINFO /home/moresec/conf/zoneinfo.zip

EXPOSE 8001
COPY --from=builder /workdir/{{.ModPrefix}}{{.Name}}/{{.ModPrefix}}{{.Name}}  /app/{{.ModPrefix}}{{.Name}}
