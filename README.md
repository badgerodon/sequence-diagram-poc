# Sequence Diagram POC

Proof of Concept Sequence Diagram Generator

## build

```bash
go generate ./...
env GOOS=js GOARCH=wasm go build -v -o dist/main.wasm .
```

## local web server

```bash
(cd dist && goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))')
```
