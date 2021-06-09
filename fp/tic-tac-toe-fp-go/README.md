# Tic-tac-toe

A console 3x3 Tic-tac-toe for 2 friends.

[![GoReportCard example](https://goreportcard.com/badge/github.com/dibikhin/tic-tac-toe-go)](https://goreportcard.com/report/github.com/dibikhin/tic-tac-toe-go) [![Maintainability](https://api.codeclimate.com/v1/badges/229dc45729c3983e99a9/maintainability)](https://codeclimate.com/github/dibikhin/tic-tac-toe-go/maintainability) [![example branch parameter](https://github.com/dibikhin/tic-tac-toe-go/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/dibikhin/tic-tac-toe-go/actions/workflows/go.yml)

## How to

Get someone and play locally in your terminal using keyboard only. Cannot play w/ computer yet so can play with yourself at worst :) See how to run below.

## Getting Started

### Prerequisites
- Install [Go](https://golang.org/doc/install) (tested on go1.15.7 linux/amd64)
- Zero dependencies

### Installing
Not needed, runs as is, just clone:
```
$ cd my_projects/
$ git clone https://github.com/dibikhin/tic-tac-toe-go.git
```

## Running the tests
```
$ cd tic-tac-toe-go/
$ make test
...
>PASS
>ok      tictactoe/game  0.010s
```

NOTE: The tests play the game itself too, scroll to see.

## Running
```
$ cd tic-tac-toe-go/
$ make run
> Hey! This is 3x3 Tic-tac-toe for 2 friends :)
>
> X   X
> O X O
> X   O
>
> Press 'x' or 'o' to choose mark for Player 1:
```

NOTE: Hit `ctrl+c` to exit.

## Building & Starting
```
$ cd tic-tac-toe-go/
$ make
>Testing...
>...
>Building...
>...
$ make start
```

## Internals

### Project Structure
- `/game` — The game package
- `go.mod`
- `main.go` — Entry point
- `Makefile`

### Overview
- The UI is CLI
- The 3x3 size is hardcoded
- No timeouts for turns
- Dirty input tolerant

### Technical
- The code has excessive amount of comments for educational purposes
- The game is one active app (no client/server)
  - Simple but structured
  - Zero deps
  - No patterns overkill
- A basic DI is under the hood for auto-tests (naive, no too smart DI)
  - a simple IoC in `setup.go`, can stub/mock the user input strategy
  - an inner DI in the game loop (game ctx in `loop.go`)
- Well-tested
  - no mocks (behavior), just stubs (state)
  - ~90% code coverage
  - pure and atomic fns mainly (no IO tests)
  - NOTE: The tests play the game itself too. See in the end after expanding the `Test` section of [the Github Actions job 'build'](https://github.com/dibikhin/tic-tac-toe-go/runs/2290602609?check_suite_focus=true)
- Paradigm: OOP + FP principles, SRP enforced
- Functional programming: a lot of pure fns; IO extracted but not isolated

## Authors
- [Roman Dibikhin](https://github.com/dibikhin)

## License
This project is licensed under the MIT License — see the [LICENSE](./LICENSE) file for details.

## Acknowledgments
Thanks to:
- [A Tour of Go](https://tour.golang.org/welcome/1) — For the idea
- [Tic-tac-toe](https://en.wikipedia.org/wiki/Tic-tac-toe) — A lot of insights about the game
