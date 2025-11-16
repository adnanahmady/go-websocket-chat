# go-websocket-chat
A simple real-time chat server using Go WebSockets.

## Run Application

You can run the application using bellow command.

```shell
make run
```

## Frontend

You can check the application behaviour using prepared
small frontend application provided in `localhost:5000/chat`
url.

## Test

You can run tests using this command

```shell
make test
```

## Project Structure

You can take a brief look at the whole projects structure in below

```shell
.
├── config
│   ├── loader.go
│   └── structs.go
├── config.yml # default config variable values
├── go.mod
├── go.sum
├── internal # application modules
│   ├── integration # tests
│   │   ├── join_room_test.go
│   │   ├── leave_room_test.go
│   │   └── send_message_test.go
│   ├── presentation
│   │   ├── v1_handlers.go
│   │   └── v1_routes.go
│   ├── wire_gen.go
│   └── wire.go
├── LICENSE
├── main.go
├── Makefile # prepared commands helper
├── pkg # shared packages
│   ├── applog
│   │   └── logger.go
│   ├── request
│   │   ├── contexts.go
│   │   ├── middlewares.go
│   │   └── server.go
│   ├── test
│   │   └── setup.go
│   └── websocket
│       ├── client.go
│       ├── hub.go
│       └── message.go
├── README.md
└── static # front end app, written with alpinejs
    ├── alpine.js.3.15.1.min.js
    ├── index.html
    ├── style.css
    └── tailwind@4.js
```