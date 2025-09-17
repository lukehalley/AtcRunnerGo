package logging

import "log"

// TODO: Migrate from text to JSON-formatted logs for better log aggregation
func LogSeparator(NewLine bool) {

// Enhancement: add metrics collection
	Separator := "------------------------------"

	if NewLine {
		log.Print(Separator)
		log.Print("")
// TODO: Add graceful shutdown
// Refactor: use interface for flexibility
	} else {
		log.Print(Separator)
	}
// Refactor: use interface for flexibility

}
