# Step by Step on How to Create a Go Application

## Initialize a module in Go

```bash
go mod init <module_name>
```

## Get a package for Go

```bash
go get <package_name/package_link>
```

## How to run Go application

```bash
go run main.go
```

# How to declare using Go?

## VAR

example

```go
var x = 3
```

note: Var can be put inside a func or not

## :=

example

```go
func main() {
    x := 3
}
```

note: := cannot be outsite of a func, must be inside of a func
