package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"sort"
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

func ReadFile(filename string) []string {
	// Read the file into a variable
	// Open the file
	dat, err := os.Open(filename)
	check(err)
	// Close the file and make sure it is closed after everything is done
	defer dat.Close()
	// We will be reading the file word by word or here number by number


	scanner := bufio.NewScanner(dat)
	scanner.Split(bufio.ScanWords)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	return text

}

func main() {
	fmt.Println("Hello, World!")
	
	text := ReadFile("input.txt")

	var list1 []int
	var list2 []int

	for i := 0; i < len(text); i+=2 {
		num1, _ := strconv.Atoi(text[i])
		num2, _ := strconv.Atoi(text[i+1])
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}
	// Sort both the slices
	sort.Sort(sort.IntSlice(list1))
	sort.Sort(sort.IntSlice(list2))

	// Calculate the distance for part 1
	// var dis int
	// for i:=0; i<len(list1); i++ {
	// 	dis += abs(list1[i], list2[i])
	// }

	// fmt.Println(dis) = 1834060

	// Calculate the answer for part 2
	ans := make(map[int]int)

	for i:=0; i<len(list1); i++ {
		key := list1[i]
		count := 0
		for j:=0; j<len(list2); j++ {
			if key == list2[j] {
				count++
			}
		}
		ans[key] = count
	}
	var dis int
	for i := range ans {
		dis += (i * ans[i])
	}

	fmt.Println(dis)
}

func abs(x int, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}