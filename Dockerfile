# build
FROM golang:1.10 AS build
WORKDIR /opt/go-mail

ADD main /opt/go-mail
ADD /toml/qq-mail.toml /opt/go-mail
ADD /toml/company-email.toml /opt/go-mail

CMD  ["./main","-config=./qq-mail.toml"]

