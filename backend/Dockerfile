FROM golang:1.22

RUN apt-get update && apt-get install -y git

WORKDIR /code

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
