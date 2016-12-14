package explorepackages

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func SimpleLog() {
	fmt.Println("***************Show Logs in the Standard Output***************")
	// Display a simple log in the stdout
	log.Println("Simple log using log.Println() function")

	// This does not get displayed as 'ioutil.Discard' ignores printing this log
	tracelog := log.New(ioutil.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	tracelog.Println("This is discarded as its a part of trace")

	// Prints the error log on stdout with a full path (due to the use of flag 'Llongfile' from log package )
	errorlog := log.New(os.Stderr, "ERROR: ", log.Ltime|log.Ldate|log.Llongfile)
	errorlog.Println("Your code is breaking -- its a error")

	//Prints information on stdout with just a file name and line number
	infolog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	infolog.Println("This is a information")

	//Prints Warning to stdout with the file name and line number
	warninglog := log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	warninglog.Println("This is a Warning")
}

func LogToFile() {
	fmt.Println("***************Save Logs in the File alogfile.log***************")
	// Create file to save the log
	file, e := os.Create("alogfile.log")
	if e != nil {
		fmt.Printf("Meeli")
	}

	// Writes a TRACE to the file
	tracelog := log.New(file, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	tracelog.Println("This is discarded as its a part of trace")

	// Writes a ERROR to the file with the full path and line number
	errorlog := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
	errorlog.Println("Your code is breaking -- its a error")

	// Writes a INFORMATION to the file with the file name and line number
	infolog := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	infolog.Println("This is a information")

	// Writes a WORNING to the file with file name and line number
	warninglog := log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	warninglog.Println("This is a Warning")

	// Writes a simple log in the file
	log.SetOutput(os.Stdout)
	log.Println("Simple Log : with flags ")
	log.SetFlags(0) //-- can be used to remove all the flags and only show the error message
	log.Println("Simple Log : log.SetFlags(0) can be used to remove all the flags and only show the error message ")
}

func MultiWrite() {
	fmt.Println("***************Show and Save Logs***************")
	// Creates a new file to show the log
	file, e := os.Create("error.log")
	if e != nil {
		fmt.Printf("Meeli")
	}
	// Creates multiwrite with file and stdout.
	multiWriter := io.MultiWriter(file, os.Stdout)

	// Creates a customized log object and use the same to print error
	multiWriterLog := log.New(multiWriter, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	multiWriterLog.Println("This error is shown in file and stdout")
}
