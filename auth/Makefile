BINARY=engine
all: test clean run

test : 
		go test ./...

build:
		go build -o ${BINARY} main.go

clean:
		@echo "cleaning built apps from local storage..."
		@if [ -f ${BINARY} ] ; then rm -f ${BINARY} ; fi

docker:
		docker build -t commo .

run:
		docker-compose up --build -d

stop:
		docker-compose down

.PHONY: test clean build docker run stop