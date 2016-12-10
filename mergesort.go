package main

import "time"
import "math/rand"
import "fmt"
import "sort"
import "runtime"

const arrayLen = 100000000
const stThreshold = 100000

var arr [arrayLen]int

func mergeSort(a []int, ch chan bool) {
    mid := len(a)/2

    if len(a) < stThreshold {
        st_mergeSort(a[:mid])
        st_mergeSort(a[mid:])
        merge(a, mid)
        ch <- true
        return
    }

    ch1 := make(chan bool)
    ch2 := make(chan bool)

    go mergeSort(a[:mid], ch1)
    go mergeSort(a[mid:], ch2)

    x := <-ch1
    y := <-ch2

    merge(a, mid)

    ch <- x && y
}


func st_mergeSort(a []int) {
    mid := len(a)/2
    if len(a) > 1 {
        mid = len(a)/2
        st_mergeSort(a[:mid])
        st_mergeSort(a[mid:])
        merge(a, mid)
    }
}


func merge(a []int, mid int) {
    l1 := len(a[:mid])
    l2 := len(a[mid:])

    var buffer1 = make([]int, l1)
    var buffer2 = make([]int, l2)

    copy(buffer1, a[:mid])
    copy(buffer2, a[mid:])

    x, y := 0, 0
    p := 0

    for x < l1 && y < l2 {
        if buffer1[x] <= buffer2[y] {
            a[p] = buffer1[x]
            x++
        } else {
            a[p] = buffer2[y]
            y++
        }
        p++
    }

    if x < l1 {
        copy(a[p:], buffer1[x:])
    }

    if y < l2 {
        copy(a[p:], buffer2[y:])
    }
}


func main() {
    runtime.GOMAXPROCS(8)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var a []int = arr[0:]

    sum := 0

    numLoop := 1

    for loop := 0; loop < numLoop; loop++ {

        for i := range a{
            a[i] = r1.Intn(arrayLen)
        }

        start := int(time.Now().UnixNano())
        c := make(chan bool)
        go mergeSort(a, c)
        b := <-c
        end := int(time.Now().UnixNano())

        if b {
            sum += end - start
        }

        if (!sort.IntsAreSorted(a)) {
            fmt.Println("Wrong !")
            for i := 0; i < len(a); i++ {
                fmt.Println(a[i])
            }
        }
    }

    fmt.Println(float64(sum/numLoop)/1000000000)
}
