package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputFilePath := "./InputOutputFiles/input.txt"
	outputFilePath := "./InputOutputFiles/output.txt"

	inFile, _ := os.Open(inputFilePath)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	var outputBuffer bytes.Buffer
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Printf("Line is: %s \n", line)

		inputs := strings.Split(line, "|")
		fmt.Printf("Employee ID: %s, Email: %s \n", inputs[0], inputs[1])

		c, d, e := GetSolutionForInput(inputs[0], inputs[1])
		//fmt.Printf("%s | %s | %s | %s | %s \n", inputs[0], inputs[1], c, d, e)

		outputBuffer.WriteString(inputs[0])
		outputBuffer.WriteString(" | ")
		outputBuffer.WriteString(inputs[1])
		outputBuffer.WriteString(" | ")
		outputBuffer.WriteString(c)
		outputBuffer.WriteString(" | ")
		outputBuffer.WriteString(d)
		outputBuffer.WriteString(" | ")
		outputBuffer.WriteString(e)
		outputBuffer.WriteString("\n")
	}

	fmt.Print("\n Meeli Vyas Output\n")
	fmt.Print(outputBuffer.String())

	var outputFile, _ = os.Create(outputFilePath)
	writer := bufio.NewWriter(outputFile)
	defer outputFile.Close()

	fmt.Fprintln(writer, outputBuffer.String())
	writer.Flush()
}

func GetSolutionForInput(a string, b string) (string, string, string) {
	return "firstName", "lastName", "phone"
}
