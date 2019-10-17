FROM node:12.12.0-buster

WORKDIR /usr/src/app

COPY package.json yarn.lock ./
#RUN apk add --no-cache --virtual .gyp python make g++
RUN yarn --network-timeout 1000000
#RUN apk del .gyp
COPY . ./

RUN ["yarn", "build"]
RUN ["yarn", "global", "add", "serve"]
CMD ["serve", "-s", "build"]
