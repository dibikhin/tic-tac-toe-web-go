# Tic-tac-toe for web Go

A web version of 3x3 Tic-tac-toe on Go: server + terminal client.

## Other versions
- [Funcional programming, in terminal](https://github.com/dibikhin/tic-tac-toe-cli-fp-go)
- [Object-oriented programming, in terminal](https://github.com/dibikhin/tic-tac-toe-cli-oop-go)

## How to
Start server and clients, get a friend. Play in terminal using keyboard only. See how to run below.

## Getting Started

### Prerequisites
- Install [Go](https://golang.org/doc/install) (tested on go1.17.2 linux/amd64)
- gRPC

### Installing
```
$ cd my_projects/
$ git clone https://github.com/dibikhin/tic-tac-toe-web-go.git
$ cd ./tic-tac-toe-web-go/
$ go mod download
$ cp ./cmd/client/example.env ./cmd/client/.env
$ cp ./cmd/server/example.env ./cmd/server/.env
```

### Testing
Run and look how tests play themselves.

```
$ cd tic-tac-toe-web-go/
$ make test
Testing client...
...
PASS
...
Testing server...
...
PASS
...
```

## Running
Run as is, no compilation neaded. Running locally assumed.

### Server
Open first terminal, then:
```
$ cd tic-tac-toe-web-go/
$ make serve
Server:
2022/04/19 11:15:24 Starting...
2022/04/19 11:15:24 Started
...
```

NOTE: Hit `ctrl+c` to exit.

### Client
Open second and third terminal, then in each one:
```
$ cd tic-tac-toe-web-go/
$ make connect
Client:
2022/04/19 11:17:56 Starting...
...
2022/04/19 11:17:56 Started

Welcome to 3x3 Tic-tac-toe for 2 friends :)
...
```

NOTE: Hit `ctrl+c` to exit.

## Internals

### Features
- The UI is CLI
- The 3x3 size is hardcoded
- No timeouts for turns
- Dirty input tolerant
- Server handles unlimited games
- State refresh via short-polling
- Games are stored in memory only
- Client handles loosing connection well (reconnects automatically)
- Dead game detection
- Multiple clients for the same player name
- Game survives restarting client
- No soft limit for games' count
- A first player waits a second one

## Authors
- [Roman Dibikhin](https://github.com/dibikhin)

## License
This project is licensed under the MIT License — see the [LICENSE](./LICENSE) file for details.

## Acknowledgments
Thanks to:
- [A Tour of Go](https://tour.golang.org/welcome/1) — For the idea
- [Tic-tac-toe](https://en.wikipedia.org/wiki/Tic-tac-toe) — Some insights about the game
