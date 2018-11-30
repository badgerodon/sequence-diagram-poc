# Sequence Diagram POC

Proof of Concept Sequence Diagram Generator

## generate

Install https://github.com/mna/pigeon and then:

```bash
go generate ./...
```

This is only necessary if the PEG grammar changes.

## build

```bash
env GOOS=js GOARCH=wasm go build -v -o dist/main.wasm .
```

## local web server

```bash
(cd dist && goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))')
```
