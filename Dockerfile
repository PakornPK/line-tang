FROM alpine:3.17
WORKDIR /app
COPY bin/ /app/

EXPOSE 55753

CMD [ "./main" ]
