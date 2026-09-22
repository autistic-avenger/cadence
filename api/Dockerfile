FROM ubuntu:latest

COPY . ./api/

WORKDIR /api/

RUN apt update -y
RUN apt upgrade -y
RUN apt install golang-go ca-certificates -y

EXPOSE 8000

CMD ["go","run","./cmd/server"]

