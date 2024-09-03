package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/abdulkhaliq7/GO_AWK/awk"
)

func main() {
	start := time.Now()
	file, err := os.Open("cdr.txt")
	if err != nil {
		log.Println(err)
	}

	/*
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		newAwk := awk.NewAwk(line)
		output, err := newAwk.DataSplit("|", "")
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

	*/
}
