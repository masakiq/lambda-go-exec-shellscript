## Install package

```
$  GOROOT=/your-go-root GOPATH=/your-go-path go get -t -d -v ./...
```

## Build

```
$ GOOS=linux GOARCH=amd64 go build -o hello
```

## Convert zip

```
$ zip -r handler.zip hello hello-bash
```
