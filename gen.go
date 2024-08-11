package gen

//go:generate go run ./cmd/tools/terntodotnet/main.go
//go:generate sqlc generate -f ./internal/store/pgstore/sqlc.yaml
