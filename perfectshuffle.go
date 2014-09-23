package main

import "fmt"

func main() {
    var i int
    for fmt.Scanln(&i); i > 0; fmt.Scanln(&i) {
        cards1 := make([]string, i)
        for j := 0; j < (i / 2 + i % 2); j++ {
            fmt.Scanln(&cards1[j])
        }
        cards2 := make([]string, i)
        for j := 0; j < i / 2; j++ {
            fmt.Scanln(&cards2[j])
        }
        cards1Ind := 0
        cards2Ind := 0
        for j := 0; j < i; j++ {
            if j % 2 == 0 {
                fmt.Println(cards1[cards1Ind])
                cards1Ind++
            } else {
                fmt.Println(cards2[cards2Ind])
                cards2Ind++
            }
        }
    }
}