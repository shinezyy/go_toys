package main

import "time"
import "math/rand"
import "fmt"
import "sort"
import "runtime"

const arrayLen = 30000
const stThreshold = 10000

var arr [arrayLen]int


func st_rankSort(a []int) []int {
    ranks := make([]int, len(a))
    ret := make([]int, len(a))
    for i := range(a) {
        rank := 0
        for j := range(a) {
            if a[i] > a[j] {
                rank++
            }
            if a[i] == a[j] && i > j {
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

        ret := st_rankSort(a)

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
