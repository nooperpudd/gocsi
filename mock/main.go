package main

import (
	"context"

	"github.com/nooperpudd/gocsi"
	"github.com/nooperpudd/gocsi/mock/provider"
	"github.com/nooperpudd/gocsi/mock/service"
)

// main is ignored when this package is built as a go plug-in
func main() {
	gocsi.Run(
		context.Background(),
		service.Name,
		"A Mock Container Storage Interface (CSI) Storage Plug-in (SP)",
		"",
		provider.New())
}
