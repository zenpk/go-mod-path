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
root_dir/
├── go.mod
└── some_folder/
    └── parent/
        └── go.mod/
            └── your_package/
                ├── code.go <--- use gmp in this file
                └── go.mod
```

Then the two functions will work like this

```go
// same as gmp.GetPath(0)
gmp.GetNearestPath() // root_dir/some_folder/parent/your_package/

gmp.GetPath(1) // root_dir/some_folder/parent/
gmp.GetPath(2) // root_dir/
```

