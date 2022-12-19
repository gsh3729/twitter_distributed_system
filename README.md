# Twitter v2.0

This is a minimal Twitter app that is written on Golang and uses etcd as storage system and communicates across services using grpc.

Technical specifications:

1. Frontend using gin-gonic - A golang framework for developing web apps easily. We use their sessions feature to login users and store sessions on the browser as cookies. No states are stored in the frontend, and communication to backend is done using gRPC.
2. Backend Micro Services - Authentication, Tweets, and Followers. These services are used to provide the features of our system. They are golang services which implement gRPC servers.
3. etcd is our storage tool. etcd is a distributed, reliable key-value store for the most critical data of a distributed system. It uses RAFT algorithm to reach consensus and maintain availability and concurrency across multiple storage replicas.

---

> ### Supported features:
> 1. User Signup
> 2. User Signin
> 3. List Users on the platform
> 4. Follow and Unfollow other users
> 5. Post tweets
> 6. View tweets that you or people you follow have posted
> 7. Session based user login

---

## Instructions To Run App
### Setup

Install [Golang](https://go.dev/doc/install)

### Install and Start etcd

Run the install_etcd.sh script from root directory.
```bash
./install_etcd.sh
```
Run the runraft.sh script from root directory.
```bash
./runraft.sh
```
This uses [etcd](https://github.com/etcd-io/etcd) and [goreman](https://github.com/mattn/goreman).
etcd is installed into /tmp/ directory. goreman uses the Procfile in the root directory to spin up 3 instances of etcd in a single cluster. etcd cluster becomes available for usage after all 3 replicas start.

### Start the Backend Server

From the root directory:
```bash
    ./runbackend.sh
```

### Run Tests

Run test cases for backend using:
```bash
    cd ./backend/authbackend
    go test 

    cd ./backend/follow
    go test 

    cd ./backend/tweet
    go test 

```

### Start the Frontend

From root of project.  

```bash
    go run ./web/server.go
```

> The application should be available at https://ide8000.anubis-lms.io/ or http://localhost:8080/ if running locally. 


