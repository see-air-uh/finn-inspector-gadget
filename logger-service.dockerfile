FROM alpine:latest 

RUN mkdir /app 

COPY LoggerServiceApp /app

CMD ["/app/LoggerServiceApp"]