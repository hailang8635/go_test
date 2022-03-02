package main

import (
    "fmt"
    "time"
)


func main() {
    testIntToRoman()
}
func testLuckyNumbers() {
    fmt.Println(time.Now())
    matrix := [][]int{
      {3,7,8},
      {9,11,13},
      {15,16,17}}
    fmt.Println( luckyNumbers(matrix) )

    matrix2 := [][]int{
      {1,10,4,2},
      {9,3,8,7},
      {15,16,17,12}}
    fmt.Println( luckyNumbers(matrix2) )

    matrix3 := [][]int{
       {7,8},
       {1,2}}
    fmt.Println( luckyNumbers(matrix3) )

    // TODO
    matrix4 := [][]int{
        {3,6},
        {7,1},
        {5,2},
        {4,8}}
    fmt.Println( luckyNumbers(matrix4) )


    fmt.Println(time.Now())
}

// 在同一行的所有元素中最小
// 在同一列的所有元素中最大
//
func luckyNumbers(matrix [][]int) []int {
    n := len(matrix)
    m := len(matrix[0])

    var luckNum [] int

    var indexRowMap = make(map[int]int, n)
    var indexColumnMap = make(map[int]int, n)

    for i:=0; i<n; i++ {
        var min = matrix[i][0]
        indexRowMap[i] = 0

        // 一行里最小的，记录坐标
        var j int
        for j = 0; j < m; j++ {
            if matrix[i][j] < min {
                min = matrix[i][j]
                indexRowMap[i] = j
            }
        }
    }


    // 一列里最大，记录坐标
    var j int
    for j = 0; j < m; j++ {
        var max = matrix[0][j]
        indexColumnMap[j] = 0

        for i:=0; i<n; i++ {
            if matrix[i][j] > max {
                max = matrix[i][j]
                indexColumnMap[j] = i
            }
        }
    }

    // rk: rowIndex, rv: columnIndex
    for rk,rv := range indexRowMap {
        // ck: columnIndex, cv: rowIndex
        for ck,cv := range indexColumnMap {
            if rk == cv && rv == ck {
                luckNum = append(luckNum, matrix[rk][rv])
            }
        }

    }

    //fmt.Println(luckNum)
    return luckNum
}


func testIntToRoman() {
    fmt.Println(time.Now())
    fmt.Println(intToRoman(romanToInt("III")))
    fmt.Println(intToRoman(romanToInt("IV")))
    fmt.Println(intToRoman(romanToInt("IX")))
    fmt.Println(intToRoman(romanToInt("LVIII")))
    fmt.Println(intToRoman(romanToInt("MCMXCIV")))
    fmt.Println(time.Now())
}

func intToRoman(num int) string {
    var result = ""

    cMap := make(map[string]int)
    // I， V， X， L，C，D 和 M
    c1 := "I"
    c5 := "V"
    c10 := "X"
    c50 := "L"
    c100 := "C"
    c500 := "D"
    c1000 := "M"

    num1 := 1
    num5 := 5
    num10 := 10
    num50 := 50
    num100 := 100
    num500 := 500
    num1000 := 1000

    cMap[c1] = num1
    cMap[c5] = num5
    cMap[c10] = num10
    cMap[c50] = num50
    cMap[c100] = num100
    cMap[c500] = num500
    cMap[c1000] = num1000

    i1000 := num / 1000
    i100 := num % 1000 / 100
    i10 := num % 100 / 10
    i1 := num % 10

    result += getRepeat(i1000, c1000)

    if i100 <= 3 {
        result += getRepeat(i100, c100)
    } else if i100 == 4 {
        result += "CD"
    } else if i100 == 9 {
        result += "CM"
    } else {
        result += c500 + getRepeat(i100-5, c100)
    }

    if i10 <= 3 {
        result += getRepeat(i10, c10)
    } else if i10 == 4 {
        result += "XL"
    } else if i10 == 9 {
        result += "XC"
    } else {
        result += c50 + getRepeat(i10-5, c10)
    }

    if i1 <= 3 {
        result += getRepeat(i1, c1)
    } else if i1 == 4 {
        result += "IV"
    } else if i1 == 9 {
        result += "IX"
    } else {
        result += c5 + getRepeat(i1-5, c1)
    }

    return result
}

func getRepeat(num int, c string) string {
    result := ""
    for i:=0; i<num; i++ {
        result += c
    }
    return result
}

func testRoman() {
    fmt.Println(time.Now())
    fmt.Println(romanToInt("III"))
    fmt.Println(romanToInt("IV"))
    fmt.Println(romanToInt("IX"))
    fmt.Println(romanToInt("LVIII"))
    fmt.Println(romanToInt("MCMXCIV"))
    fmt.Println(time.Now())
}

func romanToInt(s string) int {
    var result = 0
    var length = len(s)
    if length < 1 {
        return result
    }


    cMap := make(map[string]int)
    // I， V， X， L，C，D 和 M
    c1    := "I"
    c5    := "V"
    c10   := "X"
    c50   := "L"
    c100  := "C"
    c500  := "D"
    c1000 := "M"

    num1    := 1
    num5    := 5
    num10   := 10
    num50   := 50
    num100  := 100
    num500  := 500
    num1000 := 1000

    cMap[c1]    = num1
    cMap[c5]    = num5
    cMap[c10]   = num10
    cMap[c50]   = num50
    cMap[c100]  = num100
    cMap[c500]  = num500
    cMap[c1000] = num1000

    //fmt.Println(cMap)

    if length == 1 {
        return cMap[s]
    }

    tmpNode    := ""

    for k, v := range s {
        str := string(v)
        //fmt.Println(k, str, cMap[str] )

        if k < len(s)-1 && cMap[str] < cMap[string(s[k+1])] {
            tmpNode = str
            continue
        } else if tmpNode != "" {
            if tmpNode == c1 && (str == c5 || str == c10) {
                result += cMap[str] - cMap[tmpNode]
            } else if tmpNode == c10 && (str == c50 || str == c100) {
                result += cMap[str] - cMap[tmpNode]
            } else if tmpNode == c100 && (str == c500 || str == c1000) {
                result += cMap[str] - cMap[tmpNode]
            }
            tmpNode = ""
        } else {
            result += cMap[str]
        }
    }

    return result

}

func testLetterCombinations() {
    fmt.Println(time.Now())

    fmt.Println(letterCombinations(""))
    fmt.Println(letterCombinations("2"))
    fmt.Println(letterCombinations("23"))
    fmt.Println(letterCombinations("7793"))
    fmt.Println(letterCombinations("777"))
    fmt.Println(time.Now())

}

//
func letterCombinations(digits string) []string {
    var result []string
    if (len(digits)) > 4 {
        return result
    }


    var digitsMap = make(map[string][]string)
    digitsMap["2"] = []string{"a", "b", "c"}
    digitsMap["3"] = []string{"d", "e", "f"}
    digitsMap["4"] = []string{"g", "h", "i"}
    digitsMap["5"] = []string{"j", "k", "l"}
    digitsMap["6"] = []string{"m", "n", "o"}
    digitsMap["7"] = []string{"p", "q", "r", "s"}
    digitsMap["8"] = []string{"t", "u", "v"}
    digitsMap["9"] = []string{"w", "x", "y", "z"}

    length := len(digits)

    arrAll := make([][]string, 4)
    for i:=0; i<length; i++ {
        arrAll[i] = digitsMap[string(digits[i])]
    }
    //fmt.Println(arrAll)

    //var i,j,m,n int
    for i:=0; length >=1 && i<len(arrAll[0]); i++ {
        if length == 1 {
            // ==1
            result = append(result, arrAll[0][i])
            //fmt.Println(" --> 1", result)
        } else if length > 1 {
            for j:=0; j < len(arrAll[1]); j++ {
                if length == 2 {
                    // ==2
                    result = append(result, arrAll[0][i] + arrAll[1][j])
                    //fmt.Println(" --> 2", result)
                } else if length > 2 {
                    for m:=0; m < len(arrAll[2]); m++ {
                        if length == 3 {
                            // ==3
                            result = append(result, arrAll[0][i] + arrAll[1][j] + arrAll[2][m])
                            //fmt.Println(" --> 3", result)
                        } else if length > 3 {
                            for n := 0; n < len(arrAll[3]); n++ {
                                if length == 4 {
                                    // ==4
                                    result = append(result, arrAll[0][i] + arrAll[1][j] + arrAll[2][m] + arrAll[3][n])
                                    //fmt.Println(" --> 4", result)
                                } else {
                                    fmt.Println("暂不支持4层以上的组合")
                                }
                            }
                        }
                    }
                }
            }
        }
    }

    //fmt.Println(result, "\n")
    return result
}
