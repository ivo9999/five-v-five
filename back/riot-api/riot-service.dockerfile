FROM alpine:latest

RUN mkdir /app

COPY riotApp /app

CMD [ "/app/riotApp" ]
