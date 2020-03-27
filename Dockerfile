FROM golang:1.14.1-alpine3.11

# RUN apt-get -y update 
# RUN apt-get -y install curl
# RUN apt-get -y install iputils-ping
# RUN apt-get -y install mc
# RUN apt-get -y install nodejs
# RUN apt-get -y install sqlite3 

WORKDIR /app
COPY . .

# RUN cp -r dump/. uploads
# RUN ls uploads
# RUN go version
RUN go build

CMD ./message-scheduler -port=8088 -env=docker

EXPOSE 8088