FROM ubuntu:24.04

ENV TZ='Asia/Tokyo'

RUN apt update && apt install -y wget

WORKDIR /tmp
RUN wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb && apt install -y ./google-chrome-stable_current_amd64.deb

ADD ./moneyforward /usr/local/bin/moneyforward

CMD ["/usr/local/bin/moneyforward"]

