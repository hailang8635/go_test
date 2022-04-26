package main

import (
    "fmt"
    "github.com/dhowden/tag"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"
)
func mv() {
    path, _ := filepath.Split("Q:/音频类/test/许巍/曾经的你(吉他版).mp3")
    err := os.MkdirAll(path, os.ModePerm)
    fmt.Println(err)
    err = os.Rename("Q:\\音频类\\test\\00000.mp3", "Q:/音频类/test/许巍/曾经的你(吉他版).mp3")
    fmt.Println(err)

}

func main() {
    //mv()
    mvMp3("Q:\\音频类\\(.mp3) MP3 声音文件")
}

func mvMp3(parentDir string) {

    _, err := os.Stat(parentDir)
    if err != nil {
        log.Println("目录异常", parentDir)
        log.Fatal(err)
    }
    //TestPath("Q:/音频类/test/")
    files, err := FixFileName(parentDir)

    if err != nil {
        log.Fatal("FixFileName失败", err)
    }

    if len(files) == 0 {
        log.Println("队列为空，无需操作")
        return
    }

    for _, fileElement := range files {
        log.Println(fileElement[0], " --> ", fileElement[1])
    }

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
                fmt.Println("创建文件夹出错", basePath, err)
            }

            err = os.Rename(fileElement[0], fileElement[1])
            if err != nil {
                log.Println("Rename error", err)
            }
        }
    }

}

func FixFileName(parentDir string) (files [][]string, err error) {
    files = make([][]string, 0, 100)

    filepath.Walk(parentDir,
        func (fileName string, file os.FileInfo, err error) error {

            if file.IsDir() {
                fmt.Println("目录", fileName)
                return nil
            } else {
            }
            fileRead, _ := os.Open(fileName)
            metaInfo, err := tag.ReadFrom(fileRead)

            defer fileRead.Close()

            if err != nil {
                log.Println("tag.ReadFrom error ", err)
                return nil
            }

            //fmt.Println(metaInfo.Title(), file.Name())
            ext := filepath.Ext(file.Name())
            if metaInfo.Title() + ext == file.Name() {
                fmt.Println("\t 无需重命名 ", fileName)
                return nil
            } else {
            }

            midPath := []string{parentDir}

            artist := metaInfo.Artist()
            if metaInfo.Artist() == "" {
                artist = file.Name()
            }

            midPath = append(midPath, artist)

            if metaInfo.Album() != "" {
                midPath = append(midPath, metaInfo.Album())
            }

            if metaInfo.Title() != "" {
                midPath = append(midPath, metaInfo.Title() + ext)
            }

            newName := strings.Join(midPath, "/")
            fileElement := []string{fileName, newName}

            //log.Println(fileName, " ---> ", newName)
            files = append(files, fileElement)

            return nil
        })

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