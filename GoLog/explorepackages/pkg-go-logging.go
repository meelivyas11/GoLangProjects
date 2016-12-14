package explorepackages

import (
	"fmt"
	"os"

	logging "github.com/op/go-logging"
)

//go-logging leveling are: DEBUG < INFO < NOTICE < WARNING < ERROR < CRITICAL
// When a specific level is selected, logs related to that level and higher are shown
// using go-logging package
func LogLevel() {
	fmt.Println("***************Show Leveled Logs***************")
	var log = logging.MustGetLogger("example")

	// Setup Level - Shows warnings and higher levels
	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.NOTICE, "")

	logging.SetBackend(backend1Leveled)
	log.Debugf("Leveled debug log without format")
	log.Infof("Leveled information log without format")
	log.Noticef("Leveled notice log without format")
	log.Warningf("Leveled warning log without format")
	log.Errorf("Leveled err log without format")
	log.Criticalf("Leveled crit log without format")
}

func LogFormat() {
	fmt.Println("*************** Show Formated Logs***************")
	var log = logging.MustGetLogger("example")
	// Setup Formater
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)
	var format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
	backend2Formatter := logging.NewBackendFormatter(backend2, format)

	logging.SetBackend(backend2Formatter)
	log.Debugf("Formatted debug")
	log.Infof("Formatted information")
	log.Noticef("Formatted notice")
	log.Warningf("Formatted warning")
	log.Errorf("Formatted err")
	log.Criticalf("Formatted crit")
}

func LogLevelWithFormat() {
	fmt.Println("***************Show Formated and Leveled Logs ***************")
	var log = logging.MustGetLogger("example")

	// Create a formatter
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)
	var format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`)
	backend2Formatter := logging.NewBackendFormatter(backend2, format)

	// Set Levels on the formatter
	backend3FormatLevel := logging.AddModuleLevel(backend2Formatter)
	backend3FormatLevel.SetLevel(logging.WARNING, "")

	logging.SetBackend(backend3FormatLevel)
	log.Debugf("Leveled formatted debug")
	log.Infof("Leveled formatted information")
	log.Noticef("Leveled formatted notice")
	log.Warningf("Leveled formatted warning")
	log.Errorf("Leveled formatted err")
	log.Criticalf("Leveled formatted crit")
}
