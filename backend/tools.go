//go:build tools
// +build tools

package main

import (
	_ "github.com/dmarkham/enumer"
	_ "github.com/pressly/goose/example/migrations-go"
	_ "github.com/pressly/goose/v3/cmd/goose"
)
