FROM uhub.service.ucloud.cn/entropypool/alpine:3.16

WORKDIR /app

COPY commission-manager /app

ENTRYPOINT ["/app/commission-manager"]

EXPOSE 50340 50341 50342 50343

CMD ["run"]
