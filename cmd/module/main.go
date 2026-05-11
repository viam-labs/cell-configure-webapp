// Package main is the entrypoint binary for the cell-configure-webapp
// Viam module. The module ships a single_machine application (defined
// in meta.json) and registers no Go-side models — the binary exists
// only because Viam modules require an entrypoint.
package main

import (
	"go.viam.com/rdk/module"
)

func main() {
	module.ModularMain()
}
