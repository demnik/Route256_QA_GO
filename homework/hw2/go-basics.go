// Basic tasks for golang
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите предложения разделенные точкой с пробелом:")
	inputString := readInputString()

	//какие-то долгие разборки с кириллицей.
	//hello golang. i'm not crying, the cyrillic alphabet just got into my eyes

	fmt.Println(capitalizeFirstLetterInSentences(inputString))

	fmt.Println("Введите массив чисел с плавающей точкой, каждый элемент которого разделён пробелом:")
	inputSliceFloat64 := readInputSliceFloat64()

	//1.0 2.0 3.0 4.0 5.0
	//1651.0 315.0 656.0

	fmt.Println(average(inputSliceFloat64))
	fmt.Println(differenceBetweenMaxAndMinSliceValue(inputSliceFloat64))
}

func capitalizeFirstLetterInSentences(inputString string) string {
	var outString = ""
	var lastSymbol = ""

	for i, symbol := range inputString {
		if i == 0 {
			outString += strings.ToUpper(string(symbol))
			lastSymbol = string(symbol)
			continue
		}

		if (lastSymbol == "." || lastSymbol == "!" || lastSymbol == "?") && string(symbol) != " " {
			outString += strings.ToUpper(string(symbol))
			lastSymbol = string(symbol)
			continue
		}

		outString += string(symbol)

		if string(symbol) != " " {
			lastSymbol = string(symbol)
		}
	}

	if lastSymbol != "." && lastSymbol != "!" && lastSymbol != "?" {
		outString += "."
	}

	return outString
}

func average(slice []float64) float64 {
	var sum float64

	for _, num := range slice {
		sum += num
	}

	return sum / float64(len(slice))
}

func differenceBetweenMaxAndMinSliceValue(slice []float64) float64 {
	sort.Float64s(slice)
	return slice[len(slice)-1] - slice[0]
}

func readInputString() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func readInputSliceFloat64() []float64 {
	inputString := readInputString()
	var inSliceString = strings.Split(inputString, " ")
	outSlice := make([]float64, len(inSliceString))

	for i, num := range inSliceString {
		outSlice[i], _ = strconv.ParseFloat(num, 64)
	}

	return outSlice
}
