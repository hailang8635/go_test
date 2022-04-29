package main

import (
    "bufio"
    "bytes"
    "encoding/hex"
    "fmt"
    "github.com/dhowden/tag"
    "golang.org/x/text/encoding/simplifiedchinese"
    "golang.org/x/text/transform"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"
    "unicode/utf8"
)

func mv() {
    path, _ := filepath.Split("Q:/音频类/test/许巍/曾经的你(吉他版).mp3")
    err := os.MkdirAll(path, os.ModePerm)
    fmt.Println(err)
    err = os.Rename("Q:/音频类/test/00000.mp3", "Q:/音频类/test/许巍/曾经的你(吉他版).mp3")
    fmt.Println(err)

}

func main() {
    //mv()
    //mvMp3("Q:/音频类/(.mp3) MP3 声音文件")

    //TestChar()
    //
    //mvMp3("Q:/音频类", "Q:/dest")
    mvMp3("Q:/dest2", "Q:/dest")
}

func mvMp3(parentDir string, destDir string) {

    _, err := os.Stat(parentDir)
    if err != nil {
        log.Println("目录异常", parentDir)
        log.Fatal(err)
    }
    //TestPath("Q:/音频类/test/")
    files, err := FixFileName(parentDir, destDir)

    if err != nil {
        log.Fatal("FixFileName失败", err)
    }

    if len(files) == 0 {
        log.Println("队列为空，无需操作")
        return
    }

    log.Printf("已发现文件 %d \n", len(files))


    logContent := ""
    for _, fileElement := range files {
        logContent += fileElement[0] + " --> " + fileElement[1] + "\n"
        //log.Println(fileElement[0], " --> ", fileElement[1])
    }
    writeLog("fileListLog.txt", logContent)

    fmt.Println("是否重命名(移动)文件 Y/N ?")
    var text string
    fmt.Scan(&text)

    if text == "Y" {
        fmt.Println("开始重命名... ")
        for _, fileElement := range files {
            log.Println(fileElement[0] + " --> " + fileElement[1])

            basePath, _ := filepath.Split(fileElement[1])
            err = os.MkdirAll(basePath, os.ModePerm)
            if err != nil {
                log.Fatal("创建文件夹出错", basePath, err)
            }

            err = os.Rename(fileElement[0], fileElement[1])
            if err != nil {
                log.Fatal("Rename error", err)
            }
        }
    }

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

func FixFileName(parentDir string, destDir string) (files [][]string, err error) {
    files = make([][]string, 0, 100)
    count := 0
    noTagCount := 0
    succCount := 0

    filepath.Walk(parentDir,
        func (fileName string, file os.FileInfo, err error) error {
            if file.IsDir() {
                fmt.Println("目录", fileName)
                return nil
            } else {
                count++
            }
            fileRead, _ := os.Open(fileName)
            metaInfo, err := tag.ReadFrom(fileRead)

            defer fileRead.Close()

            if err != nil {
                //log.Println("tag.ReadFrom error ", err)
                noTagCount++
                return nil
            }

            //fmt.Println(metaInfo.Title(), file.Name())
            ext := filepath.Ext(file.Name())
            fileNameForPath := ""

            if metaInfo.Title() == "" {
                fileNameForPath = file.Name()
            } else {
                if strings.Contains(metaInfo.Title(), file.Name()) || strings.Contains(file.Name(), metaInfo.Title()) {
                    fmt.Println("\t 无需重命名 ", fileName)
                    fileNameForPath = file.Name()
                    //return nil
                } else {
                    fileNameForPath = metaInfo.Title() + ext
                }
            }
            succCount++

            midPath := []string{""}

            artist := metaInfo.Artist()
            if strings.Trim(metaInfo.Artist(), "") == "" {
                artist = file.Name()
            }

            //fmt.Println(len(artist), len([]byte(artist)))
            //fmt.Println(strings.ToUpper(hex.EncodeToString([]byte(artist))))
            //c1,_ := GbkToUtf8Str(artist)
            //c2,_ := Utf8ToGbkStr(artist)
            //fmt.Println("-->", c1, "   ", c2, []byte(artist))

            midPath = append(midPath, artist)

            if metaInfo.Album() != "" {
                midPath = append(midPath, metaInfo.Album())
            }

            //if metaInfo.Title() != "" {
            //    midPath = append(midPath, fileNameForPath)
            //} else {
            //    midPath = append(midPath, file.Name())
            //}
             midPath = append(midPath, fileNameForPath)

            newName := destDir +
                strings.ReplaceAll(
                    strings.ReplaceAll(
                        strings.ReplaceAll(
                        strings.ReplaceAll(
                        strings.ReplaceAll(
                        strings.ReplaceAll(
                        strings.ReplaceAll(
                        strings.ReplaceAll(
                            strings.ReplaceAll(strings.Join(midPath, "/"), ":", ""),
                        " ",""),
                        "*",""),
                    "?",""),
                    "<",""),
                    ">",""),
                    "|",""),
                    "\\",""),
                "\"","")
            fileElement := []string{fileName, newName}

            //log.Println(fileName, " ---> ", newName)
            files = append(files, fileElement)

            return nil
    })
    log.Printf("总：%d 成功：%d 无标签：%d \n", count, succCount, noTagCount)

    return files, nil
}

func TestPath(path string) {
    if path == "" {
        path = "Q:/音频类/test/"
    }
    files, _ := ioutil.ReadDir(path)
    for _, file := range files {
        //fmt.Println(file.Name())

        f, _ := os.Open(path + file.Name())

        m, err := tag.ReadFrom(f)
        if err != nil {
            log.Println("error", err)
        }
        log.Println(m.Format())
        log.Println(m.Artist())
        log.Println(m.Album(), m.Album() == "")
        log.Println(m.Title())

        /*
        id3 := make([]byte,128)
        //_, _ = f.ReadAt(id3, file.Size() - 128)
        _, _ = f.ReadAt(id3, 0)

        utf8Data, _ := simplifiedchinese.GBK.NewDecoder().Bytes(id3)
        fmt.Println(string(utf8Data))
        gbkData, _ := simplifiedchinese.GBK.NewEncoder().Bytes(id3)
        fmt.Println(string(gbkData))
         */
    }

    //files2, _:= filepath.Glob(path + "*")
    //fmt.Println(files2)

}

// GBK 转 UTF-8
func GbkToUtf8Str(s string) (string, error) {
    reader := transform.NewReader(bytes.NewReader([]byte(s)), simplifiedchinese.GBK.NewDecoder())
    d, e := ioutil.ReadAll(reader)
    if e != nil {
        return "", e
    }
    //fmt.Println(string(d))
    return string(d), nil
}

// UTF-8 转 GBK
func Utf8ToGbkStr(s string) (string, error) {
    reader := transform.NewReader(bytes.NewReader([]byte(s)), simplifiedchinese.GBK.NewEncoder())
    d, e := ioutil.ReadAll(reader)
    if e != nil {
        return "", e
    }
    //fmt.Println(strings.ToUpper(hex.EncodeToString(d)))
    return string(d), nil
}



func GbkToUtf8(s []byte) ([]byte, error) {
    reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
    d, e := ioutil.ReadAll(reader)
    if e != nil {
        return nil, e
    }
    //fmt.Println(string(d))
    return d, nil
}

// UTF-8 转 GBK
func Utf8ToGbk(s []byte) ([]byte, error) {
    reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
    d, e := ioutil.ReadAll(reader)
    if e != nil {
        return nil, e
    }
    //fmt.Println(strings.ToUpper(hex.EncodeToString(d)))
    return d, nil
}

func TestChar() {

    //fmt.Println([]byte("十"))
    //fmt.Println([]byte("Ê®"))
    fmt.Println(hex.EncodeToString([]byte("Ê®")))
    //fmt.Println(hex.EncodeToString([]byte("Ê®¶þ´óÃÀÅ®")))

    //fmt.Println(hex.EncodeToString([]byte("TPE1")))
    //s := "十二大美女"

    // GBK:  CA AE
    // UTF8: E5 8D 81
    //unicode:
    //s := "十"

    // utf8: C38A  C2AE
    // C38A --> 195 138 --> U+00CA
    // C2AE --> 194 174 --> U+00AE
    s := "Ê®"
    //fmt.Printf("%s \n", strings.ToUpper(hex.EncodeToString([]byte(s))))

    r,_ := utf8.DecodeRuneInString("Ê")
    fmt.Printf("%x \n", r)
    r2,_ := utf8.DecodeRuneInString("®")
    fmt.Printf("%x \n", r2)


    g,_ := Utf8ToGbk([]byte(s))
    fmt.Println("  --> gbk  ", g, string(g))

    // u,_ := GbkToUtf8([]byte{0xCA, 0xAE})
    u,_ := GbkToUtf8([]byte{byte(r), byte(r2)})
    fmt.Println("  --> utf8  ", u, string(u))

    //GbkToUtf8()
    //Utf8ToGbk("十二大美女")
}
/**
Format() Format
FileType() FileType

Title() string
Album() string
Artist() string
AlbumArtist() string
Composer() string
Genre() string
Year() int

Track() (int, int) // Number, Total
Disc() (int, int) // Number, Total

Picture() *Picture // Artwork
Lyrics() string
Comment() string
*/
// Metadata