# Othelo game

## How to run?

```console
$ go run .
```

## How to play?

Stone placements inputs are done by command line input. Player will write `X` and `Y` coordinates, and the board will evaluate the result. It will place the stone if the move is legal, but it will not place the stone if the move turns out to be illegal.

Illegal move consists of 2 things.
- There isn't a single stone that's turned over as a repurcusion of the player's move
- The placement of the stone is off the dimension of the board.

```console
$ Enter X and Y pos of your move:
```

If you answer like so
```console
$ 3 5
```
it is equivalent for you placing your stone on coordinate (3, 5)
