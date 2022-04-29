package main

import (
    "bufio"
    "fmt"
    "path/filepath"
    "strconv"

    //"bytes"
    //"encoding/hex"
    //"github.com/dhowden/tag"
    //"golang.org/x/text/encoding/simplifiedchinese"
    //"golang.org/x/text/transform"
    //"io/ioutil"
    "log"
    "os"
    //"strings"
    //"unicode/utf8"
)


func main() {
    //mv()
    //mvMp3("Q:/音频类/(.mp3) MP3 声音文件")

    //TestChar()
    //
    //mvMp3("Q:/音频类", "Q:/dest")
    //llDir("Q:/test")
    // da----         2022/4/27     17:44               01.个人资料
    //da----         2022/4/23     14:41                02.个人照片
    //da----          2021/1/3     21:11                03.个人视频
    //da----         2022/4/24     12:47                05.ln_photo
    //da----         2022/4/27     18:22                13.百科全书
    //da----         2022/4/27     18:21                14.通识书籍
    //da----         2021/12/6     19:21                21.代码库
    //da----         2020/7/10     22:02                22.技术-书籍-讲座-资料
    //da----         2022/4/24      7:58                23.考试资料
    //da----          2022/4/1     21:38                31.software
    //da----         2021/4/18     10:59                51.业务资料
    //da----         2022/4/27     17:45                61.儿童节目
    //da----         2022/4/24      9:14                81.monitor
    //da----         2022/4/24     17:35                logs
    //da----         2020/8/18      0:03                opt
    //da----         2022/4/23      8:15                temp
    //da----          2020/7/5     13:04                wd_elements
    //da----         2020/9/26     15:25                wym
    //da----        2020/12/27     11:28                迅雷下载
    llDir("U:/")
    llDir("R:/")
    llDir("V:/")
    llDir("W:/")
}

func llDir(parentDir string) {

    _, err := os.Stat(parentDir)
    if err != nil {
        log.Println("目录异常", parentDir)
        log.Fatal(err)
    }
    //TestPath("Q:/音频类/test/")
    files, err := llFileName(parentDir)

    if err != nil {
        log.Fatal("llFileName", err)
    }

    if len(files) == 0 {
        log.Println("队列为空，无需操作")
        return
    }

    log.Printf("已发现文件 %d \n", len(files))


    logContent := ""
    for _, fileElement := range files {
        logContent += fileElement + "\n"
    }
    writeLog(parentDir + "/00_file_list_log.txt", logContent)

}

func writeLog(fileName string, logContent string) {
    fileListLog, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        fmt.Println("写日志失败", err)
    } else {
        fmt.Println("日志已写入", fileName)
    }
    defer fileListLog.Close()
    write := bufio.NewWriter(fileListLog)
    write.WriteString(logContent)
    write.Flush()
}

func llFileName(parentDir string) (files []string, err error) {
    files = make([]string, 0, 100)
    count := 0

    filepath.Walk(parentDir,
        func (fileName string, file os.FileInfo, err error) error {
            if file.IsDir() {
                fmt.Println("目录", fileName)
                return nil
            } else {
                count++
            }

            //log.Println(fileName, " ---> ", newName)
            line := strconv.FormatInt(file.ModTime().Unix(), 10) + "\t\t" + strconv.FormatInt(file.Size(), 10) + "\t\t" + file.Name() + "\t\t" + fileName
            files = append(files, line)

            return nil
        })
    log.Printf("总：%d  \n", count)

    return files, nil
}
