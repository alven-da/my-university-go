# Install packages

Execute `go mod init` and `go mod download`

# How to run locally

Execute `go run cmd/api/main.go`

And test if working properly by accessing `/health` API

Should return

```
{
  "status": "ok"
}
```

# How to run unit test

Execute `go test ./internal/package_folder`

e.g. `go test ./internal/usecase` to test all services in `usecase` folder
