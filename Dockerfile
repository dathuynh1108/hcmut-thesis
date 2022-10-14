FROM node:alpine
WORKDIR /service/app/server

COPY . .
COPY ./configs /configs


RUN npm install

RUN  apk add --no-cache bash \
        \
        && apk add --no-cache tzdata \
        \
        && cp /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime \
        \
        && echo "Asia/Ho_Chi_Minh" > /etc/timezone

RUN apk --no-cache add ca-certificates && apk --no-cache add curl

ENV CONFIG_PATH="./configs/server.toml"

CMD ["node", "cmd/server/main.js"]
