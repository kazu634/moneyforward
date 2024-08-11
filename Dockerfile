FROM cimg/go:1.22.6-browsers

RUN apk add --no-cache libc6-compat tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata

ADD ./moneyforward /usr/local/bin/moneyforward

ENTRYPOINT ["/usr/local/bin/moneyforward"]
