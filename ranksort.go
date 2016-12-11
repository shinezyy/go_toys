package main

import "time"
import "math/rand"
import "fmt"
import "sort"
import "runtime"
import "sync"

const arrayLen = 3000
const numSeg = 8

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


func getRank(a []int, L int, R int, ranks []int, wg *sync.WaitGroup){
    for i := L; i < R; i++ {
        rank := 0
        for j := range(a) {
            if a[i] >a[j] {
                rank++
            }
            if a[i] ==a[j] && i > j {
                rank++
            }
        }
        ranks[i] = rank
    }
    wg.Done()
}


func rankSort(a []int) []int {
    ranks := make([]int, len(a))
    ret := make([]int, len(a))

    segLen := arrayLen/numSeg
    cursor := 0
    var wg sync.WaitGroup

    for i := 0; i < numSeg; i++ {
        wg.Add(1)
        if i == numSeg - 1 {
            go getRank(a, cursor, arrayLen, ranks, &wg)
        } else {
            go getRank(a, cursor, (i+1)*segLen, ranks, &wg)
            cursor = (i+1)*segLen
        }
    }

    wg.Wait()

    for i := range(a) {
        ret[ranks[i]] = a[i]
    }
    return ret
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

        ret := rankSort(a)

        end := int(time.Now().UnixNano())

        mt_time += end - start
        if (!sort.IntsAreSorted(ret)) {
            fmt.Println("Wrong !")
        }

        start = int(time.Now().UnixNano())
        ret = st_rankSort(array_copy, array_copy)
        end = int(time.Now().UnixNano())
        st_time += end - start

        if (!sort.IntsAreSorted(ret)) {
            fmt.Println("Wrong !")
        }
    }

    fmt.Println("Multiple Thread Time:", float64(mt_time/numLoop)/1000000000)
    fmt.Println("Single Thread Time:", float64(st_time/numLoop)/1000000000)
    fmt.Println("Speedup:", float64(st_time)/float64(mt_time))
}
