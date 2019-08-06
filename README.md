# Hello World

Hello World is a simple CLI application written in Golang with Cobra

## Global info

- Download and install tools used by make
```
$ go run mage.go -d tools
```

- Download the dependences locked
```
$ make depend.vendor
```

- Check the licenses
```
$ make license
```

- Release the dev actions
```
$ make dev
```

## How to 

1. `go run mage.go`
2. `bin/hello-world --help`
3. 
 ```
 bin/hello-world say hello
 hello world !
 ```
