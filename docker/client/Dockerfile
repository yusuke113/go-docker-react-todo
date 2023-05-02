FROM node:16.14.2-alpine3.15

ENV TZ Asia/Tokyo

WORKDIR /var/www/src

COPY client/ /var/www/

RUN npm install

CMD ["npm", "run", "dev"]