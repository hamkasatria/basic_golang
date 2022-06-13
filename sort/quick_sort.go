package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"

	localLibSort "github.com/jfcg/sorty/v2"
)

func getGOMAXPROCS() int {
	// return runtime.GOMAXPROCS(0)
	return 1
}

func main() {
	runtime.GOMAXPROCS(getGOMAXPROCS())
	fmt.Printf("GOMAXPROCS is %d\n", getGOMAXPROCS())
	long := float64(7)
	num := generateRandomSlice(int(math.Pow(10, long)))
	// fmt.Println(num)
	TestQuickSort(num)
	TestGo(num)

}

func TestGo(rand []int) {

	before := time.Now()

	// localLibSort.SortSlice(rand)
	localLibSort.Sort(rand)

	after := time.Now()
	diff := after.Sub(before)

	// fmt.Println("Concurent \t: ", rand, "\n\a")
	fmt.Println("Time Con \t=", diff.Nanoseconds(), "==", diff.String())
}

func TestQuickSort(rand []int) {

	before := time.Now()

	QuickSort(rand)

	after := time.Now()
	diff := after.Sub(before)

	// fmt.Println(" Hasil \t: ", data, "\n\a")
	fmt.Println("Time nor \t=", diff.Nanoseconds(), "==", diff.String())
}

func QuickSort(slice []int) []int {
	length := len(slice)

	if length <= 1 {
		sliceCopy := make([]int, length)
		copy(sliceCopy, slice)
		return sliceCopy
	}

	m := slice[rand.Intn(length)]

	less := make([]int, 0, length)
	middle := make([]int, 0, length)
	more := make([]int, 0, length)

	for _, item := range slice {
		switch {
		case item < m:
			less = append(less, item)
		case item == m:
			middle = append(middle, item)
		case item > m:
			more = append(more, item)
		}
	}

	less, more = QuickSort(less), QuickSort(more)

	less = append(less, middle...)
	less = append(less, more...)

	return less
}

func generateRandomSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
