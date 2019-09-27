# Go apps 101: How to build & test them

Welcome to the Orders Go App! 

This repo contains the Orders App for the Golab `"Go apps 101: How to build & test them"` workshop. 

The app is written in Go. 

The slides accompanying this workshop are found on [talks.godoc](https://talks.godoc.org/github.com/addetz/go_apps_101_slides/golab19.slide#1)

## Setup 

### Go setup 

- Download & install the Go tools. Make sure you add the `bin` directory to your `PATH` as well. 
 Choose the installation steps for your OS at [golang.org](https://golang.org/doc/install#install)
	- You will need to install Go 1.12 or newer to run this code. 
- The default workspace directory is `$HOME/go`. If you'd like to use a different directory, you will need to setup your [GOPATH environment variable](https://github.com/golang/go/wiki/SettingGOPATH).
- Set an environment variable `$GO111MODULE=on`. 
	- This is only required for Go versions of `1.11` & `1.12`, as `1.13` has this enabled by default.
- Test your installation easily. Create a new file `main.go` in your `GOPATH` and add the code below to it. 
```go
package main

import "fmt"

func main() {
	fmt.Printf("All set and ready to Go! \n")
}
```
- Run your code with the command below and make sure you see the output `All set and ready to Go!`. 
```
$ go run main.go
```

### Repo setup 
Clone this repo into `$GOPATH/src/github.com` directory.

## Technologies used 
- Go 1.12
- [gorilla/mux](https://github.com/gorilla/mux)
- [matryer/moq](https://github.com/matryer/moq)

## Author
Adelina Simion

Software Engineer, Back End @ Deliveroo 

[@classic.addetz](https://twitter.com/classic_addetz)

[@DeliverooEng](https://twitter.com/DeliverooEng)


![Deliveroo logo](deliveroo-main.png)

![GoLab logo](golab-logo.png)
