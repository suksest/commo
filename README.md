# commo

## Description
This RESTful API provide:

- Auth
- Commodity Info (aggregate not implemented yet)

### How To Run This Project

Since the auth service already use Go Module, I recommend to put the source code in any folder but GOPATH.

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
$ git clone https://github.com/suksest/commodity.git

#move to project
$ cd commodity

# Run the application
$ make run

# Or you can run default run script
$ make all

# check if the containers are running
$ docker ps

# Execute the call
$ curl localhost:17845/auth

# Stop
$ make stop
```

### API Documentation
See [API.md](https://github.com/suksest/commodity/blob/master/API.md)