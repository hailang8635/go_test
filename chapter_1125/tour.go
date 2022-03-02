package chapter_1125

import (
    "fmt"
    "golang.org/x/tour/pic"
    "os"
    "strings"
)

func Split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func Pic(dx, dy int) [][]uint8 {
    result := make([][]uint8, dy)

    //fmt.Println(result, cap(result))

    //var i uint8 = 0
    //for resultElement := range result {
    for i := 0; i < cap(result); i++ {
        resultX := make([]uint8, dx)

        for j := range resultX {
            //resultX[j] = uint8(i * j)
            //resultX[j] = uint8(i * j)
            //resultX[j] = uint8(i * j)
            resultX[j] = uint8( (i + j) / 2 )
        }

        result[i] = resultX

        //result = append(result, resultX)
    }

    //fmt.Println(result)

    return result

}

func Show() {
    pic.Show(Pic)

    fmt.Println("")
}



func WordCount(s string) map[string]int {
    words := strings.Split(s," ")
    //fmt.Println(words[0:])

    resultMap := make(map[string]int)

    for idx := range words {
        //fmt.Println(words[idx])
        count, ok := resultMap[words[idx]]
        if ok {
            //fmt.Println(count)
            count++
            resultMap[words[idx]] = count
        } else {
            resultMap[words[idx]] = 1
            //fmt.Println(word)
        }
    }

    fmt.Println(resultMap)
    return resultMap
    // return map[string]int{"x": 1}
}


func TestPanic () {
    panic("a problem")

    _, err := os.Create("target/file")
    if err != nil {
        panic(err)
    }
}