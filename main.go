package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

func ShellSort(arr []int) []int {
	n := len(arr)
	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i
			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = temp
		}
	}
	return arr
}

func QuickSort(arr []int, low, high int) []int {
	if low < high {
		pi := partition(arr, low, high)
		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
	return arr
}

func medianOfThree(arr []int, low, mid, high int) int {
	if arr[low] > arr[mid] {
		if arr[mid] > arr[high] {
			return mid
		} else if arr[low] > arr[high] {
			return high
		} else {
			return low
		}
	} else {
		if arr[low] > arr[high] {
			return low
		} else if arr[mid] > arr[high] {
			return high
		} else {
			return mid
		}
	}
}

func partition(arr []int, low, high int) int {
	mid := (low + high) / 2
	pivotIndex := medianOfThree(arr, low, mid, high)
	arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex]

	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := make([]int, mid)
	right := make([]int, len(arr)-mid)
	copy(left, arr[:mid])
	copy(right, arr[mid:])

	left = MergeSort(left)
	right = MergeSort(right)

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(1500)
	}
	return arr
}

func copyArray(arr []int) []int {
	arrCopy := make([]int, len(arr))
	for i, v := range arr {
		arrCopy[i] = v
	}
	return arrCopy
}

func main() {
	fmt.Printf("OS: %s, Arch: %s, Go: %s, CPUs: %d\n",
		runtime.GOOS, runtime.GOARCH, runtime.Version(), runtime.NumCPU())

	sizes := []int{1000, 5000, 10000, 50000, 100000, 250000, 500000, 1000000}

	fmt.Printf("%-8s %-12s %-12s %-12s %-12s %-12s %-12s\n",
		"Size", "Shell(ms)", "Shell_S(ms)", "Quick(ms)", "Quick_S(ms)", "Merge(ms)", "Merge_S(ms)")
	fmt.Println(strings.Repeat("-", 80))

	for _, size := range sizes {
		randomArray := generateRandomArray(size)

		shellArray := copyArray(randomArray)
		start := time.Now()
		sortedShell := ShellSort(shellArray)
		timeShell := time.Since(start)

		quickArray := copyArray(randomArray)
		start = time.Now()
		sortedQuick := QuickSort(quickArray, 0, len(quickArray)-1)
		timeQuick := time.Since(start)

		mergeArray := copyArray(randomArray)
		start = time.Now()
		sortedMerge := MergeSort(mergeArray)
		timeMerge := time.Since(start)

		start = time.Now()
		ShellSort(sortedShell)
		timeShellSorted := time.Since(start)

		start = time.Now()
		QuickSort(sortedQuick, 0, len(sortedQuick)-1)
		timeQuickSorted := time.Since(start)

		start = time.Now()
		MergeSort(sortedMerge)
		timeMergeSorted := time.Since(start)

		fmt.Printf("%-8d %-12.3f %-12.3f %-12.3f %-12.3f %-12.3f %-12.3f\n",
			size,
			float64(timeShell.Nanoseconds())/1e6,
			float64(timeShellSorted.Nanoseconds())/1e6,
			float64(timeQuick.Nanoseconds())/1e6,
			float64(timeQuickSorted.Nanoseconds())/1e6,
			float64(timeMerge.Nanoseconds())/1e6,
			float64(timeMergeSorted.Nanoseconds())/1e6)
	}
}
