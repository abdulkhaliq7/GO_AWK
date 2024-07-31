package awk

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Awk(data, field string, columnNumber []string) {

	var fieldsChosen string

	for _, value := range columnNumber {

		c := value

		column, err := strconv.Atoi(c)

		if err != nil {

			log.Printf("the string Converter did not work : %v", err)
		}

		column = column - 1

		columnChose := strings.Split(data, "|")

		fieldsChosen += fmt.Sprintf("%v  ", columnChose[column])

	}

	fmt.Println(fieldsChosen)
}
