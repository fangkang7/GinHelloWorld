package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	strings := uniqBySlice([]string{"580808", "461163", "586209", "108889", "18389", "234867", "234867", "220951", "135856", "135855", "30290"})
	fmt.Println(strings)
	// "580808","461163","586209","108889","18389","234867","234867","220951","135856","135855","30290"
}

func uniqBySlice(strs []string) []string {
	uniqueStrs := make([]string, 0, len(strs))
	strMap := make(map[string]bool)
	for _,str := range strs {
		if !strMap[str] {
			strMap[str] = true
			uniqueStrs = append(uniqueStrs,str)
		}
	}

	return uniqueStrs
}

func sliceSwitchString(strs []int) string  {
	strNums := make([]string,len(strs))
	for i,num := range strs {
		strNums[i] = strconv.Itoa(num)
	}

	result := strings.Join(strNums, ",")

	return result;
}
