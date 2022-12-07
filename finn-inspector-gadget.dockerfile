FROM golang:1.19.3-alpine

WORKDIR /app

COPY . ./

RUN go mod download

RUN cd cmd/api/ && go build -o logger
# RUN go build

EXPOSE 50001

CMD [ "cmd/api/logger" ]