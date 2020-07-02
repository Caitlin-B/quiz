package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//add a command line flag, shown when you run the file with --help 
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'") //(name, value, usage)
	flag.Parse()
	
	//read the file
	file, err := os.Open(*csvFilename)

	// deal with error
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s\n", *csvFilename)) //means something went wrong
	}

	//create csv reader
	r :=  csv.NewReader(file)

	//parse csv
	lines, err := r.ReadAll()

	//deal with err
	if err!= nil {
		exit("Failed to parse the provided CSV file.")
	}
	fmt.Println(lines)
	
}


func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}