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

Game screen look like so;
```console
[Game turn 41]

Y > 7 :| | |w|w|w|w|w|w|
Y > 6 :| | | | |b|w|w|w|
Y > 5 :| | | |b|b|b|w|w|
Y > 4 :|b|b|b|b|b|b|w|w|
Y > 3 :| | | |b|w|w|b|w|
Y > 2 :| | |b|b|w|w|w|w|
Y > 1 :| | | |w|w|w|w|w|
Y > 0 :| | |w|w|w|w|w|w|
        0 1 2 3 4 5 6 7
        ^ ^ ^ ^ ^ ^ ^ ^
        X X X X X X X X

Now Turn: Black
White: 31, Black: 14
```
