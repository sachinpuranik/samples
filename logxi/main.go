// You can edit this code!
// Click here and start typing.
package main

import (
	"logxi/internal/foo"
	"os"

	"github.com/mgutz/logxi/v1"
)

// create package variable for Logger interface
// var logger log.Logger

func main() {
	// use default logger
	who := "mario"
	log.Info("Hello", "who", who)

	// create a logger with a unique identifier which
	// can be enabled from environment variables
	logger = log.New("pkg")

	// specify a writer, use NewConcurrentWriter if it is not concurrent
	// safe
	modelLogger = log.NewLogger(log.NewConcurrentWriter(os.Stdout), "models")

	db, err := sql.Open("postgres", "dbname=testdb")
	if err != nil {
		modelLogger.Error("Could not open database", "err", err)
	}

	fruit := "apple"
	languages := []string{"go", "javascript"}
	if log.IsDebug() {
		// use key-value pairs after message
		logger.Debug("OK", "fruit", fruit, "languages", languages)
	}
}
