# About
Poking at Exercise 2.2 on p. 44 of the Go Programming Language book and learning about Go Modules in the process.

This is a working Go Modules example that is *outside* the `$GOPATH`.

```
~/scripts/unit_conversion$ tree
.
├── go.mod
├── mainapp
│   └── main.go
└── unit_conversion
    └── unit_conversion.go

2 directories, 3 files
```

Notice that the `go.mod` file is in the root/parent directory, and below that are both the `mainapp/` dir and the `unit_conversion` dir, which houses the `unit_conversion` Go Package.

To create the `go.mod` file, this was run:

```
$ go mod init fake.com/unit_conversion
```

Notice that `main.go` is in it's own folder, a folder that is a sibling of the `unit_conversion` package. Also, note that there are no `go.mod` files in either `mainapp/`, nor `unit_conversion/`. 

The `import` in `main.go` looks like this:

```
"fake.com/unit_conversion/unit_conversion"
```

To call functions / vars inside `unit_conversion.go` from `main.go`, we use the pacakge name then a dot then the variable / function name.

```
unit_conversion.FeetToMetersPretty(1)
```

# Using Go Module Resources
https://medium.com/faun/golang-package-management-using-go-modules-d3c929900114  
https://medium.com/@adiach3nko/package-management-with-go-modules-the-pragmatic-guide-c831b4eaaf31  
