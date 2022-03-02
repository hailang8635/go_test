package main

import (
    "fmt"
    "go_test_1122/chapter_0209"
    "go_test_1122/chapter_1125"
    "go_test_1122/chapter_chan"
    "go_test_1122/stun"
    "time"
)

func main() {
    //testChapter1125()
    //test_chan()
    //testDay1()
    testDay2()
}

func testDay2() {
    fmt.Println(time.Now())
    //fmt.Println(chapter_0209.ThreeSum(nums))
    //fmt.Println(chapter_0209.MergeTwoLists(nums))
}
func testDay1() {
    //nums := []int{2,7,11,15}
    //var targetNum int = 9
    //nums := []int{2,5,5,11}
    //var targetNum int = 10
    //fmt.Println(chapter_0209.TwoSum(nums, targetNum))

    //nums := []int{-1,0,1,2,-1,-4}
    nums := []int{1,2,-2,-1}

    start := time.Now().UnixMilli()
    fmt.Println(time.Now())
    //fmt.Println(chapter_0209.ThreeSum(nums))
    fmt.Println(len(chapter_0209.ThreeSum(nums)))
    //fmt.Println(len(chapter_0209.ThreeSumBak(nums)))
    fmt.Println(time.Now())
    end := time.Now().UnixMilli()
    fmt.Println(end - start)

}

func test_chan() {
    chapter_chan.Test_main()
}

func testChapter1125() {
    //chapter_1125.TestPanic()

    // chapter_1125.WordCount("Package strings implements simple functions to manipulate strings implements simple functions to manipulate strings implements simple functions to manipulate strings to manipulate strings to manipulate strings")

    // chapter_1125.TestDeferInFile()

    // chapter_1125.TestJSON()

    chapter_1125.Show()

    //var result1, result2 = chapter_1125.Split(17)
    //fmt.Println(result1)
    //fmt.Println(result2)
}

func testChapterHttp() {
    //chapter_http.Test_go_http1()
    //chapter_http.Test_go_http_router()

    //chapter_http.Test_go_http2()
    //chapter_http.Test_go_http3()


    //chapter_http.TestNewRouterMux()
    //
    //chapter_http.TestMysql()
    //chapter_http.Test_sessions()


    stun.TestStun()

}


