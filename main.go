package main

import (
	"github.com/Drax-1/tandem/internal/cmd"
	"github.com/Drax-1/tandem/internal/logging"
)

func main() {
	defer logging.RecoverPanic("main", func() {
		logging.ErrorPersist("Application terminated due to unhandled panic")
	})

	cmd.Execute()
}
