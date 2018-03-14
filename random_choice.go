package randomchoice 
import (
    "math/rand"
    "fmt"
    "time"
)

func RandomChoice(count int, choiceCount int) []int{
    if  count < choiceCount{
        choiceCount = count
    }

    rand.Seed(time.Now().UnixNano())  
    intSlice := make([]int, count)
    for i:=0;i<count;i++{
        intSlice[i] = i        
    }

    idx := 0
    for i:=0;i<choiceCount;i++{
        idx = rand.Int() % count + i
        // swap
        swap(intSlice, i, idx)
        count--

    }
    return intSlice[0:choiceCount]
}

func swap(s []int, i int, j int){
    s[i], s[j] = s[j], s[i]
}

