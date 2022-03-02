package chapter_0209

import (
    "fmt"
    "sort"
)

func ThreeSum(nums []int) [][]int {
    //fmt.Println(nums)
    //var resultKey1, resultKey2, resultKey3 int
    //var result1, result2, result3 int

    result := make([][]int, 100000)
    var index = 0

    resultMap := make(map[TripeInt]bool)

    targetSumMap := make(map[int]int)
    for key1, num1 := range nums {
        targetSumMap[num1] = key1
    }

    for key2, num2 := range nums {
        nums3 := nums[key2+1:]
        for key3, num3 := range nums3 {
            //fmt.Println("[", key1, key2, key3, "] --> ", num1, num2, num3)
            _, ok2 := targetSumMap[0- (num2 + num3)]
            if ok2 && (targetSumMap[0- (num2 + num3)] != key2 && targetSumMap[0- (num2 + num3)] != key3+key2+1) {

                resultChild := [] int{0- (num2 + num3), num2, num3}
                sort.Ints(resultChild)

                tripeInt := TripeInt{resultChild[0], resultChild[1], resultChild[2]}
                _, ok := resultMap[tripeInt]
                if !ok {
                    resultMap[tripeInt] = true

                    //fmt.Println(resultChild)
                    //fmt.Println("resultChild tripeInt -->", 0 - (num2 + num3), num2, num3, tripeInt)
                    //append(result, resultChild)
                    result[index] = resultChild
                    index++
                    //break
                }
            }
        }
    }

    //fmt.Println()

    //result := [][] int{{resultKey1, resultKey2, resultKey3}}
    return result[0:index]
}


func ThreeSumBak(nums []int) [][]int {
    //fmt.Println(nums)
    //var resultKey1, resultKey2, resultKey3 int
    //var result1, result2, result3 int

    result := make([][]int, 30000)
    var index = 0

    resultMap := make(map[TripeInt]bool)

    for key1, num1 := range nums {
        nums2 := nums[key1+1:]
        for key2, num2 := range nums2 {
            nums3 := nums2[key2+1:]
            for _, num3 := range nums3 {
                //fmt.Println("[", key1, key2, key3, "] --> ", num1, num2, num3)
                if num1 + num2 + num3 == 0 {
                    //fmt.Println("[", resultKey1, resultKey2, resultKey3, "] --> ", result1, result2, result3)

                    resultChild := [] int{num1, num2, num3}
                    sort.Ints(resultChild)

                    tripeInt := TripeInt{resultChild[0], resultChild[1], resultChild[2]}
                    _, ok := resultMap[tripeInt]
                    if !ok {
                        resultMap[tripeInt] = true

                        //fmt.Println(resultChild)
                        //fmt.Println("resultChild tripeInt -->", result1, result2, result3, tripeInt)
                        //append(result, resultChild)
                        result[index] = resultChild
                        index++
                        //break
                    }
                }
            }
        }
    }

    //fmt.Println()

    //result := [][] int{{resultKey1, resultKey2, resultKey3}}
    return result[0:index]
}

type TripeInt struct {
    left, middle, right int
}



func TwoSum(nums []int, target int) []int {
    var flag bool
    var resultKey1, resultKey2 int
    //fmt.Println(resultKey1, resultKey2)
    for key1, num1 := range nums {
        fmt.Println(nums[key1+1:])
        for key2, num2 := range nums[key1+1:] {
            if num1 + num2 == target {
                flag = true
                resultKey1 = key1
                resultKey2 = key1+1+key2

                //fmt.Println("[", resultKey1, resultKey2, "]", num1, num2)
                break
            }
        }

        if flag {
            break
        }
    }

    result := [] int{resultKey1, resultKey2}
    return result
}
