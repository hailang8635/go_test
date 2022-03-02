package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println(time.Now())
    nums := make([]int, 100)
    nums[10] = 10
    nums[30] = 30
    nums[50] = 50
    fmt.Println(removeDuplicates(nums))
    fmt.Println()
    fmt.Println()


    nums2 := []int{0, 1, 2, 2, 2, 3, 4, 5, 6}
    fmt.Println(removeDuplicates(nums2))
    fmt.Println()
    fmt.Println()

    nums3 := []int{1,1,2}
    fmt.Println(removeDuplicates(nums3))
    fmt.Println()
    fmt.Println()

    //fmt.Println(nums2[:3], nums2[3+1:])
    fmt.Println(time.Now())

}

func removeDuplicates(nums []int) int {
    for i := 1; i<len(nums); i++ {
        if nums[i-1] == nums[i] {
            //fmt.Println(nums[curr], nums[i])
            nums = append(nums[:i], nums[i+1:]...)
            i--
        } else {
        }
    }

    //fmt.Println("return : ", len(nums), nums)
    return len(nums)
}
func removeDuplicatesBak2(nums []int) int {
    curr := 0

    var head = []int{nums[0]}

    for i := 1; i<len(nums); i++ {
        if nums[curr] == nums[i] {
            //fmt.Println(nums[curr], nums[i])
        } else {
            nums[curr] = nums[i]
            curr++
        }
    }

    nums = append(head, nums[0:curr]...)

    fmt.Println("return : ", len(nums), nums)
    return len(nums)
}

func removeDuplicatesBak(nums []int) int {

    numsUniqueMap := make(map[int]bool)
    for i := 0; i<len(nums); i++ {
        _, existsFlag := numsUniqueMap[nums[i]]
        //fmt.Println(i, nums[:i], nums[i+1:], nums, nums[i] )
        if existsFlag {
            //delete
            nums = append(nums[:i], nums[i+1:]...)
            i--
        } else {
            numsUniqueMap[nums[i]] = true
        }
    }
    //fmt.Println("return : ", len(nums), nums)
    return len(nums)
}
