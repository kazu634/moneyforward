FROM alpine:latest

RUN apk add --no-cache libc6-compat tzdata chromium && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata

ADD ./moneyforward /usr/local/bin/moneyforward

CMD ["/usr/local/bin/moneyforward"]
