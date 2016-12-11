package main

import "time"
import "math/rand"
import "fmt"
import "sort"
import "runtime"
import "sync"

const arrayLen = 10000000
const stThreshold = 100000

var arr [arrayLen]int


func mergeSort(a []int, wg *sync.WaitGroup) {
    mid := len(a)/2

    if len(a) < stThreshold {
        st_mergeSort(a[:mid])
        st_mergeSort(a[mid:])
        merge(a, mid)
        wg.Done()
        return
    }


    var wg2 sync.WaitGroup
    wg2.Add(2)
    go mergeSort(a[:mid], &wg2)
    go mergeSort(a[mid:], &wg2)
    wg2.Wait()


    merge(a, mid)

    wg.Done()
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

        var wg sync.WaitGroup
        wg.Add(1)
        go mergeSort(a, &wg)
        wg.Wait()

        end := int(time.Now().UnixNano())

        mt_time += end - start
        if (!sort.IntsAreSorted(a)) {
            fmt.Println("Wrong !")
        }

        start = int(time.Now().UnixNano())
        st_mergeSort(array_copy)
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
