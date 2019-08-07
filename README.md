# Hello World

Hello World is a simple CLI application written in Golang with Cobra.

## Pre-requisits

Install Go in 1.12 version minimum.

## Global info

- Download and install tools used by make
```
$ go run mage.go -d tools
```

- Download the dependences locked
```
$ go run mage.go go:deps
```

- Release the dev actions
```
$ go run mage.go
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

### Run the app

***Arguments***

The hello-world app need some parameters :

|                 ARG					 | Pattern			       | Description														    |
|:--------------------------------------:|:-----------------------:|------------------------------------------------------------------------|
| help		                             | 		                   | application help			|
| \-\-config							 | start --config path/to/file            | path to the config file											            |
| config new		                     | 		                   | generate a new config			|
| client		                             | 		                   | query the gRPC server			|
| server		                             | 		                   | gRPC serverwith  services		|
| \-\-feature-gates		                     | key=value		                   | activate or deactivate a feature			|


#### Run gRPC server

```
$ bin/hello-world server grpc
```

#### Run gRPC client and call Greeter service

```
$ bin/hello-world client greeter sayHello -s 127.0.0.1:5555 <<< '{"name": "me"}'
{
  "message": "hello me"
}%
```
