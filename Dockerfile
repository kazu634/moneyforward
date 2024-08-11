FROM cimg/go:1.22.6-browsers

ADD ./moneyforward /usr/local/bin/moneyforward

ENTRYPOINT ["/usr/local/bin/moneyforward"]
