package main

import "time"
import "math/rand"
import "fmt"

var threshold int = 10

var arr [100000000]int


func st_qsort(a []int) []int {
	if len(a) < 2 { return a }

	left, right := 0, len(a) - 1

	// Pick a pivot
	pivotIndex := 0

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	st_qsort(a[:left])
	st_qsort(a[left + 1:])


	return a
}


func qsort(a []int, ch chan bool) {
	if len(a) < 3000000 {
		st_qsort(a)
		ch <- true
		return
	}

	left, right := 0, len(a) - 1

	// Pick a pivot
	pivotIndex := 0

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	ch1 := make(chan bool)
	ch2 := make(chan bool)

	// Go down the rabbit hole
	go qsort(a[:left], ch1)
	go qsort(a[left + 1:], ch2)

	x := <-ch1
	y := <-ch2

	ch <- x && y
}


func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var a []int = arr[0:]


	for i := range a{
		a[i] = r1.Intn(100000000)
	}

	start := int(time.Now().UnixNano())
    c := make(chan bool)
	go qsort(a, c)
    b := <-c
	end := int(time.Now().UnixNano())

    if b {
        fmt.Println(end - start)
    }

	l := len(a) - 1
	for i := 0; i < l; i++{
		if a[i] > a[i+1] {
			fmt.Println("Wrong !")
			break
		}
	}
}
