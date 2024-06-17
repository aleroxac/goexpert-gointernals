# preemption: go-1.13 vs go-1.14


## how to run
``` shell
docker run -v ${PWD}:/test -w /test golang:1.13 go run main.go
docker run -v ${PWD}:/test -w /test golang:1.14 go run main.go
```
