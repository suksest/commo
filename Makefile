BINARY=engine
all: run

run:
		docker-compose up --build -d

stop:
		docker-compose down

.PHONY: clean build docker run stop