# Builder
FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /app

COPY . .

RUN make build

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app

WORKDIR /app

# ONLY FOR DEVELOPMENT USE
# Set environment variable explicitly is not recommended
# while running on production environment, --
# use secret management tools
ENV SERVICE_PORT=17845 \
    DB_HOST=commo_mysql \
    DB_PORT=3306 \
    DB_NAME=commo_db \
    DB_USERNAME=commo \
    DB_PASSWORD=commo \
    JWT_SECRET=secret

EXPOSE 17845

COPY --from=builder /app/engine /app

CMD /app/engine
