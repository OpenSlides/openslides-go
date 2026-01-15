// Package metagen contains go definitions automatically created from the meta
// repository.
package metagen

//go:generate  sh -c "go run gen_field_def/main.go > field_def.go"
