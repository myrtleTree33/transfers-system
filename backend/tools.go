//go:build tools
// +build tools

package main

import (
	_ "github.com/dmarkham/enumer"
	_ "github.com/pressly/goose/v3/cmd/goose"
	_ "github.com/vektra/mockery/v2"
)
