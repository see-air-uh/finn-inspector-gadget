FROM alpine:latest 

RUN mkdir /app 

COPY inspectorGadget /app

CMD ["/app/inspectorGadget"]