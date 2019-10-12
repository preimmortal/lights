FROM node:8.16.2-stretch-slim

WORKDIR /usr/src/app

COPY package.json yarn.lock ./
RUN yarn
COPY . ./

CMD ["yarn", "build"]
