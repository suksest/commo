BINARY=engine
all: test clean docker run

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

.PHONY: clean build docker run stop