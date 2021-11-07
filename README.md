# Generic data structures for Go

With the upcoming release of 1.18, it's now possible to build a library of generic data structures in Go. The purpose of this library is to experiment with the implementation of generics in Go, and to build a set of production-ready, high performance data structures that take advantage of the new language features. 

I expect that parts of this library may become obsolete or less ideal over time. With each new version of Go, I expect to release a new major version of this library that: 
1) Removes any code that is now covered by the standard library
2) Updates the library to better match conventions established by the new Go version
3) Updates the library to use any new stdlib additions or language features

## How to use this

Currently, since Go 1.18 has not been released, the best way to try it out is to use `gotip`, which will manage downloading the latest (unreleased) version of `go`, and works as a CLI replacement for the `go` command.

```bash
go install golang.org/dl/gotip@latest
gotip download
gotip test ./...
```

## Data strucures

### Vec

### BinaryHeap

### LinkedList

## Other goodies

### CompareFunc

### Iterator


