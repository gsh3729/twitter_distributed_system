# Twitter v2.0

This is a minimal Twitter app that is written on Golang and uses etcd as storage system and communicates across services using grpc.

Technical specifications:

1. Frontend using gin-gonic - A golang framework for developing web apps easily. We use their sessions feature to login users and store sessions on the browser as cookies. No states are stored in the frontend, and communication to backend is done using gRPC.
2. Backend Micro Services - Authentication, Tweets, and Followers. These services are used to provide the features of our system. They are golang services which implement gRPC servers.
3. etcd is our storage tool. etcd is a distributed, reliable key-value store for the most critical data of a distributed system. 

---

> ### Currently supported features:
> 1. Create an account using Username, Name, and Password
> 2. Login with Username and Password
> 3. Follow and Unfollow other users
> 4. Post Tweets
> 5. View own Tweets
> 6. View TImeline with Tweets from people you follow
> 7. Multiple users login using sessions


---

## Instructions To Run App

The steps need to be followed in order.

### Setup

Firstly, [Install Go](https://go.dev/doc/install)

Then, install Goreman 
``` 
go install github.com/mattn/goreman@latest 
``` 

Then start from root of project and ensure go dependencies are satisfied 
``` 
    cd ./web
    go mod download 
``` 

### Start Raft

Start from root of project
```bash
    cd ./raft/src/go.etcd.io/etcd/contrib/raftexample

    go build -o raftexample

    goreman start
```

### Run Tests
Start from root of project. 

```bash
    cd ./web/users
    go test 

    cd ./web/tweets
    go test 

    cd ./web/authentication
    go test 

```
### Start the Web Server
Start from root of project. 

```bash
    go run ./web/web.go
```

### Start the Backend Services
Start from root of project.  

```bash
    go run ./web/server.go
```

> The application should be available at https://ide8000.anubis-lms.io/ or http://localhost:8080/ if running locally. 


