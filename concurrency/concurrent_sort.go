package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func sortSlice(wg *sync.WaitGroup, slice []int) {
	fmt.Println(slice)
	sort.Ints(slice)
	wg.Done()
}

func mergeSortedSlices(slice1 []int, slice2 []int) []int {
	result := make([]int, len(slice1) + len(slice2))

	var i, j, k int

	for i < len(slice1) && j < len(slice2) {
		if slice1[i] < slice2[j] {
			result[k] = slice1[i]
			i++
		} else {
			result[k] = slice2[j]
			j++
		}
		k++
	}

	for i < len(slice1) {
		result[k] = slice1[i]
		i++
		k++
	}

	for j < len(slice2) {
		result[k] = slice2[j]
		j++
		k++
	}

	return result
}

func main() {
	// 1. Consume user input

	fmt.Println("Please type a sequence of integers separated by space")

	in := bufio.NewReader(os.Stdin)
	stringInput, err := in.ReadString('\n')

	if err != nil {
		fmt.Println("Error while consuming input: " + err.Error())
		return
	}

	stringInput = strings.TrimSuffix(stringInput, "\n")

	tmp := strings.Split(stringInput, " ")

	slice := make([]int, len(tmp))

	for i, raw := range tmp {
		slice[i], err = strconv.Atoi(raw)

		if err != nil {
			fmt.Println("Error while parsing input: " + err.Error())
			return
		}
	}

    // 2. Sort slice with 4 workers

    var wg sync.WaitGroup

	var blockSize = len(slice) / 4
	var leftOver = len(slice) % 4

	fmt.Println(blockSize)

    for i := 0; i < 4; i++ {
    	var realBlockSize = blockSize
    	if i == 3 {
    		realBlockSize += leftOver
		}
		wg.Add(1)
    	go sortSlice(&wg, slice[blockSize * i : blockSize * i + realBlockSize])
	}

	wg.Wait()

	// 3. Merge sorted slice into a single one

	half1 := mergeSortedSlices(slice[0 : blockSize], slice[blockSize : 2 * blockSize])
	half2 := mergeSortedSlices(slice[2 * blockSize : 3 * blockSize], slice[3 * blockSize : 4 * blockSize + leftOver])
	sortedSlice := mergeSortedSlices(half1, half2)

	fmt.Println("Sorted slice:", sortedSlice)
}
