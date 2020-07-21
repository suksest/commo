all: test install run

test:
	npm test

install: 
	npm install

docker:
	docker build -t commo .

run:
	docker-compose up --build -d

stop:
	docker-compose down

.PHONY: all test install docker run stop