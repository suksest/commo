# commo

## Description
This RESTful API provide

- Auth
- Commodity Info (not implemented yet)

### How To Run This Project

Since the project already use Go Module, I recommend to put the source code in any folder but GOPATH.

#### Run the Testing

```bash
$ make test
```

#### Run the Applications
Here is the steps to run it with `docker-compose`

```bash
#move to directory
$ cd workspace

# Clone into YOUR $GOPATH/src
$ git clone https://github.com/suksest/commo.git

#move to project
$ cd commo

# Build the docker image first
$ make docker

# Run the application
$ make run

# check if the containers are running
$ docker ps

# Execute the call
$ curl localhost:17845/auth

# Stop
$ make stop
```

### API Documentation