FROM node:lts-alpine

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /app
COPY . .
RUN make install

ENV JWT_SECRET=secret

EXPOSE 3000
CMD [ "npm", "start" ]
