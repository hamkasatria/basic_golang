package main

import (
	"fmt"
	"sort"
)

func main() {

	// Basic()
	// MapData()
	// Sort()
	// fmt.Println(maxFunc())
	ShadMain()
}

func Basic() {
	number := []int{0, 1, 2, 3, 4, 5}
	newNumb := number[3:]

	// with interface
	fmt.Println("number 0-2 = ", GetHalfSlice(number[:2]))
	fmt.Println("number 2-~ = ", GetHalfSlice(number[2:]))
	fmt.Println("number 2-~ = ", GetHalfSlice(number[2:4]))

	//with variadic
	Variadic("Auto", "This Is Number", number...)
	CheckSlice(number[:2])
	CheckSlice(number[3:4])

	// fmt.Println("Total =", SumAll(number...))

	//
	AppendData(number, newNumb)
}

func AppendData(numb1 []int, numb2 []int) {
	for _, v := range numb2 {
		numb1 = append(numb1, v)
	}
	fmt.Println(Sort(numb1, ""))
}

func Sort(rand []int, status string) []int {
	sort.Slice(rand, func(i, j int) bool {
		return rand[i] > rand[j]
	})
	return rand
}

// with interface
func GetHalfSlice(num interface{}) interface{} {
	return num
}

// with params variadic
func Variadic(s1 string, s2 string, num ...int) {
	for _, v := range num {
		fmt.Println(s2, s1, v)
	}
}

//sumAllwith variadic
func SumAll(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

//CheckSlice
func CheckSlice(arr []int) {
	fmt.Println("panjang array = ", len(arr), "==", cap(arr))
	// for i, v := range arr {

	// }

	for i := 0; i < len(arr); i++ {
		fmt.Println(i, " = ", arr[i])
	}
}

//=============== MAP ====================//

func MapData() {
	data := map[string]string{}

	data["satu"] = "ini satu"
	data["dua"] = "ini dua"
	fmt.Println(data["satu"])
	fmt.Println(data["satu"])
	// delete(data, "dua")
	for key, v := range data {
		fmt.Println(key, "\t:", v)
	}

	da, err := data["satu"]
	fmt.Println("hasil = ", da, err)

	var data2 = make(map[string]string)
	data2["dua"] = "ini dua"
	fmt.Println(data2["dua"])

	w := "hello all"
	func(w string) {
		fmt.Println(w)
	}(w)

}

func maxFunc() int {
	number := []int{0, 1, 2, 3, 4, 5, -1, 6}
	return func(num ...int) int {
		max := 0
		for _, v := range num {
			switch {
			case v > max:
				max = v
			}
		}
		return max
	}(number...)
}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

func ShadMain() {
	data := []int{3, 5, 2, -4, 8, 11}
	count := 8
	RunCount(data, count)
}

func RunCount(ll []int, sum int) {
	data := Count(ll, func(i int, j int) bool {
		return i+j == sum
	})
	fmt.Println(data)
}

func Count(data []int, callback func(int, int) bool) [][]int {
	a := [][]int{}
	for i := range data {
		for j := i + 1; j < len(data); j++ {
			if callback(data[i], data[j]) {
				a = append(a, []int{data[i], data[j]})
			}
		}
	}
	return a
}

// func RunCount2() {
// 	list := []int{3, 5, 2, -4, 8, 11}
// 	sum := 11

// 	for i, _ := range list {
// 		data := Count2(i, list)
// 	}

// 	fmt.Println(data)
// }

// func Count2(i int, data []int) [][]int {
// 	a := [][]int{}
// 	for j := i + 1; j < len(data); j++ {
// 		if callback(data[i], data[j]) {
// 			a = append(a, []int{data[i], data[j]})
// 		}
// 	}
// 	return a
// }
