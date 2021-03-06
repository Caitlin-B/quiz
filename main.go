package main

import (
	"encoding/csv"
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

	//parse into structs
	problems := parseLines(lines)

	//get users input and check correct

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d:  %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer) //read user input
		if answer == p.a {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))

}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i , line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

//error handling fct
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}