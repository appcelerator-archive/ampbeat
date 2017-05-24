package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/appcelerator/ampbeat/beater"
)

func main() {
	err := beat.Run("ampbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
