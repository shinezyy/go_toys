package main

import "time"
import "math/rand"
import "fmt"
import "sort"
import "runtime"

const arrayLen = 400000
const stThreshold = 10000

var arr [arrayLen]int


func st_qsort(a []int) []int {
	if len(a) < 2 {
        return a
    }

	left, right := 0, len(a) - 1
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
	if len(a) < stThreshold {
        st_qsort(a)
		ch <- true
		return
	}

	left, right := 0, len(a) - 1
	pivotIndex := 0

	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	ch1 := make(chan bool)
	ch2 := make(chan bool)

	go qsort(a[:left], ch1)
	go qsort(a[left + 1:], ch2)

	x := <-ch1
	y := <-ch2

	ch <- x && y
}


func main() {
    runtime.GOMAXPROCS(8)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var a []int = arr[0:]
    array_copy := make([]int, len(a))

    mt_time := 0
    st_time :=0

    numLoop := 1

    for loop := 0; loop < numLoop; loop++ {

        for i := range a{
            a[i] = r1.Intn(arrayLen)
        }
        copy(array_copy, a)


        start := int(time.Now().UnixNano())
        c := make(chan bool)
        go qsort(a, c)
        <-c
        end := int(time.Now().UnixNano())

        mt_time += end - start
        if (!sort.IntsAreSorted(a)) {
            fmt.Println("Wrong !")
        }

        start = int(time.Now().UnixNano())
        st_qsort(array_copy)
        end = int(time.Now().UnixNano())
        st_time += end - start

        if (!sort.IntsAreSorted(array_copy)) {
            fmt.Println("Wrong !")
        }
    }

    fmt.Println("Multiple Thread Time:", float64(mt_time/numLoop)/1000000000)
    fmt.Println("Single Thread Time:", float64(st_time/numLoop)/1000000000)
    fmt.Println("Speedup:", float64(st_time)/float64(mt_time))
}
