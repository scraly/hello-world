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

- Generate configuration
```
bin/hello-world config new > helloworld.local.conf.toml
```

## How to 

### Complete build (and lint, tests ...)

`go run mage.go`

### Help

`bin/hello-world --help`

1. 
 ```
 bin/hello-world say hello
 hello world !
 ```

### Run client gRPC and call greeter service

```
$ bin/hello-world client greeter sayHello -s 127.0.0.1:5555 <<< '{}'
{
  "entity": {
    "version": "1"
  },
  "error": null
}%
```
