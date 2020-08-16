# Go Gin CRUD Sample
go & ginでCRUDなAPIを作る

## install, build, run
```
$ docker-compose up --build
```

## tree
```
.
├── Dockerfile
├── README.md
├── api
│   ├── post.go
│   └── user.go
├── config
│   └── server
│       └── server.go
├── controllers
│   ├── post_controller.go
│   └── user_controller.go
├── db
│   └── db.go
├── docker-compose.yaml
├── go.mod
├── go.sum
├── main.go
└── models
    ├── post.go
    ├── repository
    │   ├── post_psql.go
    │   └── user_psql.go
    └── user.go
```
