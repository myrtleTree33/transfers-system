FROM node:current-alpine3.17

WORKDIR /usr/src/app/api_docs

COPY package*.json ./

RUN npm ci

EXPOSE 3000

CMD ["npm", "run", "dev"]