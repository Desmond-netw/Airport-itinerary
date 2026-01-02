package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Desmond-netw/Airport-itinerary.git/utils"
)

// importing packages

// Application to pretiffer Aiport Itinerary records
func main() {
	// help flag declaration
	help := flag.Bool("h", false, "Display Usage")
	printFlag := flag.Bool("p", false, "Print output to terminal") // -p flag will print output to terminal (optional)
	flag.Parse()
	prettify := true

	// display usage if -h flag is pass in arg
	if *help {
		displayUsage()
		return
	}
	// arguments validation
	args := flag.Args()
	if len(args) != 3 {
		displayInvalidUsage() // call invalud usage display
		os.Exit(1)
	}

	inputFilePath := args[0]
	outputFilePath := args[1]
	airportCSVPath := args[2]

	// process input file
	input, ok := utils.ReadInputFile(inputFilePath)
	if !ok {
		fmt.Println("input file not found")
		return
	}

	// prettify the itinerary text
	text := utils.TrimLineBreak(input)
	if text == "" {
		fmt.Println("input file is empty")
		return
	}

	// convert airport names
	text, err := utils.ConvertNames(text, airportCSVPath)
	if err != nil {
		fmt.Println("Airport lookup malformed")
		os.Exit(1)
	}
	if text == "" {
		os.Exit(0) //nothing to write
	}

	// format timestamps
	text = utils.FormatTime(text)
	// input = text

	// prepare final output content(write everthing to file;)
	var outputContent string
	if prettify {
		outputContent = text
	} else {
		// remove ANSI codes
		outputContent = utils.RemoveANSI(text)
	}

	// write to output file
	ok = utils.WriteOutputfile(outputContent, outputFilePath)
	if !ok {
		fmt.Println("failed to write to output file")
		return
	}

	// Pint to terminal if -p flag is set
	if *printFlag {
		fmt.Println("\n" + outputContent)
	}

}

func displayUsage() {
	fmt.Println("itinerary usage: ")
	fmt.Println("go run . ./input.txt ./output.txt ./airport-lookup.csv")
	fmt.Println("\nOption:")
	fmt.Println("  -p    add -p flag to arug to print to terminal")
}
func displayInvalidUsage() {
	fmt.Println("invalid usage ")
	fmt.Println("Correct usage: go run . ./input.txt ./output.txt ./airport-lookup.csv")
	fmt.Println("\nOption:")
	fmt.Println("  -p    add -p flag to arug to print to terminal")
}
