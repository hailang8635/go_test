package main

import (
    "fmt"
    "sort"
    "time"
)

func main() {
    testSubsets()
}


// 输入： nums = [1,2,3]
// 输出：
//[
//  [3],
//  [1],
//  [2],
//  [1,2,3],
//  [1,3],
//  [2,3],
//  [1,2],
//  []
//]
//
func testSubsets() {
    //nums := []int{3,2,4,1}
    nums := []int{9,0,3,5,7}
    //nums := []int{1, 2, 3, 4}
    //nums := []int{1,2,3}
    fmt.Println(subsets(nums))
}

func subsets(nums []int) [][]int {
    length := len(nums)

    if length == 0 {
        empty := [][]int{{}}
        return empty
    }

    result := subsets(nums[:length-1])
    length2 := len(result)

    for i:=0; i<length2; i++ {
        lastChar := nums[length-1]
        newNode := append(result[i], lastChar)

        //fmt.Println("newNode ", lastChar, newNode)
        if !containsArr(result, newNode) {
            //fmt.Println("result1 ", lastChar, result)
            result = append(result, newNode)
            //fmt.Println("result2 ", lastChar, result)
        }
    }
    //fmt.Println("51 --> ", result)
    //fmt.Println()
    //fmt.Println()
    return result
}


func containsArr(arrAll [][]int, arr []int) bool {

    for _, v := range arrAll {
        //fmt.Println("arrExists ", k, ":",arrExists)
        if len(arr) == len(v) {
            b := sameArr(v, arr)
            if b {
                return b
            }
        }

    }
    return false
}

func sameArr(arrExists []int, arr []int) bool {
    sort.Ints(arr)
    sort.Ints(arrExists)
    arrExistsEqual := true
    for k2, v2 := range arr {
        if v2 != arrExists[k2] {
            arrExistsEqual = false
            break
        }
    }
    if arrExistsEqual {
        return true
    }
    return false
}



func subsetsbak(nums []int) [][]int {
    //mapAll := make(map[int[]]bool)

    var arrAll = make([][]int, 1)
    var arr = make([]int, 1)

    length := len(nums)

    for i:=0; i<length; i++ {
        for j:=i; j<length; j++ {
            arr = nums[i:j+1]
            if !containsArr(arrAll, arr) {
                arrAll = append(arrAll, arr)
            }
        }
    }
    return arrAll

}


// 输入：nums1 = [1,3], nums2 = [2]
// 输出：2.00000
func testFindMedianSortedArrays() {
    fmt.Println(time.Now())

    nums1 := []int{1,3}
    nums2 := []int{2}
    fmt.Println(findMedianSortedArrays(nums1, nums2))

    numsa := []int{1,2}
    numsb := []int{3,4}
    fmt.Println(findMedianSortedArrays(numsa, numsb))

    fmt.Println(time.Now())
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    length1 := len(nums1)
    length2 := len(nums2)
    length := length1 + length2
    arr2 := make([]int, 2)
    numsAll := make([]int, length)

    //numsAll := append(nums1, nums2...)
    //length := len(numsAll)

    if length == 0 {
        return 0.0
    }

    i := 0
    j := 0
    k := 0
    var mid = length/2+1
    //fmt.Println(mid)
    for i=0; i<mid;  {
        //fmt.Println("i -->", i)

        // 1有，2有
        // 1有，2无
        // 1无，2有
        // 1无，2无
        if j<=length1-1 {
            if k<=length2-1 {
                if nums1[j] < nums2[k] {
                    numsAll[i] = nums1[j]
                    i++
                    j++
                } else {
                    numsAll[i] = nums2[k]
                    i++
                    k++
                }
            } else {
                numsAll[i] = nums1[j]
                i++
                j++
            }
        } else {
            if k<=length2-1 {
                numsAll[i] = nums2[k]
                i++
                k++
            } else {

            }
        }

        // TODO 记录中位数
        if length % 2 == 0 {
            if i-1 == length/2-1 {
                arr2[0] = numsAll[i-1]
            }
            if i-1 == length/2 {
                arr2[1] = numsAll[i-1]

                return (float64(arr2[0]) + float64(arr2[1])) / 2
            }

        } else {
            if i-1 == length/2 {
                return float64(numsAll[i-1])
            }
        }

        //fmt.Println("i <--", i)
    }


    fmt.Println("error")
    return float64(arr2[0])
}
