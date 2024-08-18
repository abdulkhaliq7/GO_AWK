package main

import (
	"bufio"
	"fmt"
	"github/abdulkhaliq/awk/awk"
	"log"
	"os"
	"time"
)

func main() {

	// Get the current time
	start := time.Now()

	file, err := os.Open("cdr.txt")

	if err != nil {
		log.Println(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Print each line
		line := scanner.Text()
		newAwk := awk.NewAwk(line)

		output, err := newAwk.DataSplit("|", "", "35")

		if err != nil {
			log.Println("Output one has error", err)
		}

		outputTwo, err := output.DataSplit(`\`, "", "26", "70")

		if err != nil {
			log.Println("Output two has error", err)
		}

		fmt.Println(outputTwo.Data)
	}
	fmt.Println("Total time taken to complete The whole process : ", time.Since(start))

}
