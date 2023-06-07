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

# Struct Interface

## Struct

- Point dari struct adalah membangun/membuat beberapa object/properties, kurang lebihnya mirip dengan Class pada materi Javascript sebelumnya

Contoh code dari struct:

```go
type Name struct {
    FirstName string
    LastName  string
}
```

## Interface

- Point dari interface adalah membangun/membuat beberapa method yang kurang lebihnya mirip seperti konsep enkapsulasi pada Javascript, yaitu membangun dari formatter atau returning

Contoh code dari interface:

```go
type Person interface {
    getFirstName() string
}

func () getFirstName() string {
    return FirstName
}
```

# Strconv

## Apa itu strconv?

strconv merupakan sebuah package yang membantu kita untuk mengkonversikan tipe data

- Atoi : mengubah string menjadi int

# Referensi

[Slice Append Cheat Sheet](https://ueokande.github.io/go-slice-tricks/)
