# go-mod-path

A tiny module to easily get the path of `go.mod`

No more relative path, absolute path, working directory, runtime directory... Just give me the path of `go.mod` !

## Install

```shell
go get github.com/zenpk/go-mod-path
```

## Usage

Suppose your file tree looks like this

```text
root/
├── go.mod
└── folder/
    └── package/
        └── go.mod
            └── yours/
                ├── code.go <--- use gmp in this file
                └── go.mod
            └── package/
                └── go.mod
```

Then gmp (go-mod-path) will work like this

```go
// same as gmp.GetPath(0)
path, err := gmp.GetNearestPath() // root/folder/package/yours/

path, err := gmp.GetPath(1) // root/folder/package/
path, err := gmp.GetPath(2) // root/

path, err := gmp.GetFolderPath("package") // root/folder/package
```
