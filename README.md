# Gochop

> Karate Chop code challenge/kata

[Code kata](http://codekata.com/kata/kata02-karate-chop/) for binary search in Go. The kata is to implement a binary search function that halvs the search space on each iteration. The function should return the index of the target value in the sorted array, or -1 if the target value is not in the array.

## Multiple Solutions

Idea is to implement the same kata in multiple ways, to get a feel for the language and its features. Solutions are in the `cmd/main` directory in file `chops.go`. Functions are one, two, etc. The `main` function runs all the functions and prints the results.

## Running the Tests

```bash
cd cmd/main
go test -v
```

## Running the Code

```bash
cd cmd/main
go run .
```


---


