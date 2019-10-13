FROM node:8.16.2-stretch-slim

WORKDIR /usr/src/app

COPY package.json yarn.lock ./
RUN yarn
COPY . ./

RUN ["yarn", "build"]

RUN ["yarn", "global", "add", "serve"]

CMD ["serve", "-s", "build"]