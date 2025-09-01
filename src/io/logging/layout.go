package logging

import "log"

// TODO: Migrate from text to JSON-formatted logs for better log aggregation
func LogSeparator(NewLine bool) {

// Enhancement: add metrics collection
	Separator := "------------------------------"

	if NewLine {
		log.Print(Separator)
		log.Print("")
	} else {
		log.Print(Separator)
	}

}
