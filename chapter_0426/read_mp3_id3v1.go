package main

import (
    "errors"
    "fmt"
    "golang.org/x/text/encoding/simplifiedchinese"
    "io/ioutil"
    "os"
    "strings"
)
//no tag
var ErrFormat = errors.New("No TAG!")
type Mp3 struct{
    //the file attr
    Name string
    Size int64
    //the mp3 attr
    Title string
    Artist string
    Album string
    Year string
    Comment string
    Genre uint8
}
func (m *Mp3) PrintInfo(){
    fmt.Printf("%v:\n",m.Name)
    fmt.Printf("size:%v\ntitle:%v\nartist:%v\nalbum:%v\nyear:%v\ncomment:%v\n",m.Size,m.Title,m.Artist,m.Album,m.Year,m.Comment)
    fmt.Printf("genre:%v",m.Genre)
}
func (m *Mp3) GetInfo(filename string) (err error){
    id3 := make([]byte,128)
    //for read access.
    f,err := os.Open(filename)
    if err != nil{
        return
    }
    defer f.Close()
    //return FileInfo
    fileInfo,err := f.Stat()
    if err != nil{
        return
    }
    //the file size,int64
    m.Size = fileInfo.Size()
    //the file name,string. eq filename
    m.Name = fileInfo.Name()
    //if
    //_,err = f.ReadAt(id3, m.Size-128)
    _,err = f.ReadAt(id3, 0)
    if err != nil{
        return
    }
    tag := string(id3[:3])
    if tag != "TAG"{
        //err  "No ID3~"
        return ErrFormat
    }
    m.Title = strings.Trim(string(id3[3:33]),"\x00")
    m.Album = strings.Trim(string(id3[33:63]),"\x00")
    m.Artist = strings.Trim(string(id3[63:93]),"\x00")
    m.Year = string(id3[93:97])
    m.Comment = strings.Trim(string(id3[97:127]),"\x00")
    m.Genre = uint8(id3[127])
    return nil
}


//func TestId3() {
func main() {
    m := new(Mp3)

    //f,err := os.Open("Q:/音频类/(.mp3) MP3 声音文件/00000.mp3")

    findPath := "Q:/音频类/test/"
    os.Chdir(findPath)
    f,err := os.Open(findPath)
    if err != nil{
        fmt.Print(err)
    }
    defer f.Close()
    names,err := f.Readdirnames(-1)
    if err != nil{
        fmt.Print(err)
    }
    for _,name := range names{
        a := strings.Split(name,".")
        fileExt := strings.ToLower(a[len(a)-1])
        //fmt.Print(i,fileExt)
        if fileExt == "mp3"{
            err = m.GetInfo(name)
            if err != nil{
                fmt.Print(name,err)
            }
            m.PrintInfo()
        }
    }
    //fmt.Printf("%v",names)
}