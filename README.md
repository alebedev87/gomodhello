# Example of how to create a Go module in GitHub

## Module name
Go tools will query VCS of your module, so, the name needs to be a full URL to your git repository.   
```bash
$ git config remote.origin.url
https://github.com/alebedev87/gomodhello.git
$ go list -m
github.com/alebedev87/gomodhello
```

## Tags
Go modules are supposed to be tagged using SemVer with `v` prefix i.e. `v0.0.1`, `v1.0.1`.  
If `v` prefix is not specified the version is not seen as such:
```bash
$ go list -m -versions github.com/alebedev87/gomodhello
github.com/alebedev87/gomodhello v0.0.1 v0.0.2 v0.0.3
```
However still can be used:
```bash
$ go get github.com/alebedev87/gomodhello@0.0.4
go: downloading github.com/alebedev87/gomodhello v0.0.4-0.20200727134623-bb98021fc08a
go: github.com/alebedev87/gomodhello 0.0.4 => v0.0.4-0.20200727134623-bb98021fc08a
go: downloading rsc.io/quote/v3 v3.1.0
go: downloading rsc.io/sampler v1.3.0
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
```
