package chapter_chan

import (
    "fmt"
    "math/rand"
    "time"
)

func producer(header string, channel chan <- string) {
    for {
        time.Sleep(1 * time.Second)
        num := rand.Int31()
        fmt.Println("生产出：", header, num)
        channel <- fmt.Sprintf("%s(%v)", header, num)

    }
}

func customer(channel <- chan string) {
    for {
        message := <- channel
        fmt.Println("消费掉：", message)
        //message2 := <- channel
    }
}

func Test_main() {
    channel := make(chan string)

    go producer("11_cat ", channel)
    go producer("12_cat ", channel)
    //go producer("13_cat ", channel)
    //go producer("21_dog ", channel)

    go customer(channel)

    time.Sleep(10 * time.Second)

}