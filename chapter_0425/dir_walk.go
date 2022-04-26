package chapter_0425

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "os"
    "path/filepath"
    "strconv"
    "strings"
    "time"
)

func main() {
    fmt.Println("hello world")

    job := time.Now().Unix()

    files, err := WalkDir(0, "Q:/视频类/", "")
    fmt.Println()
    // fmt.Println(files, err)

    if err == nil {
        SaveFileInfo(job, files)
    } else {
        log.Fatal("获取文件列表错误", err)
    }
    //ExecSqlInMysql()
}

func WalkDir(level int, dirPath, suffix string) (files [][]string, err error) {
    splitStr := "---------------------------------------------------"
    fmt.Printf("\n遍历目录 [%d] %s %s \n", level, dirPath, splitStr)

    files = make([][]string, 0, 300000)
    suffix = strings.ToUpper(suffix)

    err = filepath.Walk(dirPath,
        func(fileNameWithPath string, file os.FileInfo, err error) error {
            if file.IsDir() {
                // 目录
                fmt.Printf("发现目录 [%d]%s \n", level+1, fileNameWithPath)
                // WalkDir(level+1, fileNameWithPath, "")

            } else {
                if suffix == "" || strings.HasSuffix(strings.ToUpper(fileNameWithPath), suffix) {
                    fmt.Printf("\t发现文件 [%d] %s %d \n", level, file.Name(), file.Size())
                    file := []string { fileNameWithPath, file.Name(), strconv.FormatInt(file.Size(), 10), ""}
                    files = append(files, file)
                } else {
                    fmt.Printf("跳过 %s \n", fileNameWithPath)
                }
            }
            return nil

    })

    return files, err
}

func SaveFileInfo(job int64, fileInfoArr [][]string) error {

    insertSql := "replace into file_find(job, file_full_name, size) values "
    //params := make([]string, 0, 100)
    params := make([]interface{}, 0, len(fileInfoArr) * 3)
    for _, fileInfo := range fileInfoArr {
        //fmt.Println(fileInfo)
        //insertSql := ` insert into file_find(job, file_full_name, size) values(?, ?, ?)`
        insertSql += "(?,?,?),"

        //param := []string{strconv.FormatInt(job, 10), fileInfo[0], fileInfo[2]}
        params = append(params, strconv.FormatInt(job, 10), fileInfo[0], fileInfo[2])
        // insertSql += strconv.FormatInt(job, 10) + "," +  fileInfo[0] + "," + fileInfo[0] + ")"
    }
    err := ExecSqlInMysql(insertSql[0:len(insertSql) - 1], params)
    if err != nil {
        return err
    }
    return nil
}

func ExecSqlInMysql(sqlStr string, params []interface{}) error {
    db, err := sql.Open("mysql", "msp:11111111@(192.168.2.43:3307)/msp?parseTime=true")

    err = db.Ping()

    //createDDL := `
    //CREATE TABLE users (
    //    id INT AUTO_INCREMENT,
    //    username TEXT NOT NULL,
    //    password TEXT NOT NULL,
    //    created_at DATETIME,
    //    PRIMARY KEY (id)
    //);`

    //_, err = db.Exec(sqlStr, params)

    stmt, _ := db.Prepare(sqlStr)
    defer stmt.Close()

    _, err = stmt.Exec(params...)
    log.Println("log println err: ", err)
    return err
}