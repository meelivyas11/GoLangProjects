package main

import (
	expPkg "GoLog/explorepackages"
	"fmt"
)

func main() {
	// Using Standard log package
	expPkg.SimpleLog()  // Show logs on standard output
	expPkg.LogToFile()  // Print logs on File (alogfile.log)
	expPkg.MultiWrite() // Prints and Save logs to the file
	fmt.Printf("\n\n")

	// Using go-logging package
	expPkg.LogLevel()
	expPkg.LogFormat()
	expPkg.LogLevelWithFormat()
	fmt.Printf("\n\n")

	// Using logrus, logrus_mate and lshook
	expPkg.LoggerConfigToShowLogs()
	//fmt.Printf("\n\n")

	//Log4go
	expPkg.Log4GoPkg()
	//fmt.Printf("\n\n")
	//log15

}
