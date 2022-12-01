# advent2022

Advent of Code 2022

## Workflow

There are some helpful functions in `util.go`. To add a new day, just run:

Unix:

```
cp -R template dayXX
```

Windows:

```
xcopy template dayXX /s /i
```

Put the input from AoC into `input.txt`. To run part 1, run

```
go run .
```

or

```
go run . part1
```

To run part 2, run

```
go run . part2
```