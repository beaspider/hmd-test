# hmd-test

## Pre Reqs

A working Go environment

## Running the code

### Linux/Mac
`export INSTAGRAM_ACCESS_TOKEN=<OAUTH_TOKEN>`

`go get github.com/beaspider/hmd-test`

`cd $GOPATH/src/github.com/beaspider/hmd-test`

`go run main.go`

### Windows
`set INSTAGRAM_ACCESS_TOKEN=<OAUTH_TOKEN>`

`go get github.com/beaspider/hmd-test`

`cd $GOPATH%/src/github.com/beaspider/hmd-test`

`go run main.go`

## Looking at `Work In Progress` with DynamoDB

### Linux/Mac
`export AWS_ACCESS_KEY_ID=<AWS_ACCESS_KEY_ID>`

`export AWS_SECRET_KEY=<AWS_SECRET_KEY>`

`cd $GOPATH/src/github.com/beaspider/hmd-test`

`go test -v`

### Windows
`set AWS_ACCESS_KEY_ID=<AWS_ACCESS_KEY_ID>`

`set AWS_SECRET_KEY=<AWS_SECRET_KEY>`

`cd $GOPATH%/src/github.com/beaspider/hmd-test`

`go test -v`


