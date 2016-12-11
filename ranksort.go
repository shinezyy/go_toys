package main

import "time"
import "math/rand"
import "fmt"
import "sort"
import "runtime"

const arrayLen = 30000
const stThreshold = 1000

var arr [arrayLen]int


func st_rankSort(a []int, target []int) []int {
    ranks := make([]int, len(a))
    ret := make([]int, len(a))
    for i := range(a) {
        rank := 0
        for j := range(target) {
            if a[i] > target[j] {
                rank++
            }
            if a[i] == target[j] && i > j {
                rank++
            }
        }
        ranks[i] = rank
    }

    for i := range(a) {
        ret[ranks[i]] = a[i]
    }
    return ret
}


func getRank(a []int, target []int, ch chan []int){
    ranks := make([]int, len(a))

    for i := range(a) {
        rank := 0
        for j := range(target) {
            if a[i] > target[j] {
                rank++
            }
            if a[i] == target[j] && i > j {
                rank++
            }
        }
        ranks[i] = rank
    }
    return ranks
}


func rankSort(a []int) []int {
    const numSeg = 2
    arrays := make([][]int, numSeg)
    ret := make([]int, len(a))

    segLen = len(arrays)/numSeg
    cursor := 0
    for i := range(arrays) {
        if i == len(arrays) - 1 {
            arrays[i] = a[cursor:]
        } else {
            arrays[i] = a[cursor: i*segLen]
            cursor = i*segLen
        }
    }

    chs := make([]chan []int, numSeg)
    for i := range(chs) {

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

        ret := st_rankSort(a, a)

        end := int(time.Now().UnixNano())

        sum += end - start

        if (!sort.IntsAreSorted(ret)) {
            fmt.Println("Wrong !")
            for i := 0; i < len(ret); i++ {
                fmt.Println(ret[i])
            }
        }
    }

    fmt.Println(float64(sum/numLoop)/1000000000)
}

/*
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
*/
