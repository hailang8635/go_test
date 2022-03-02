package chapter_http

import (
    _ "github.com/go-sql-driver/mysql"
    "time"
)
import (
    "database/sql"
    "fmt"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)
func HelloMux(writer http.ResponseWriter, request *http.Request) {
    fmt.Println("hello http mux")
}
func CreateBook(writer http.ResponseWriter, request *http.Request) {

    vars := mux.Vars(request)
    //log.Println(vars["name"])
    log.Println("vars: ", vars["title"])
    log.Println("form: ", request.FormValue("title"))
    log.Println("form: ", request.FormValue("field1"))

    fmt.Println("CreateBook ok\n\n")
    fmt.Fprintln(writer, "CreateBook ok\n\n")
}
func GetBook(writer http.ResponseWriter, request *http.Request) {
    fmt.Println("GetBook ok")
}


func TestNewRouterMux() {
    r := mux.NewRouter()

    r.HandleFunc("/hello", HelloMux)
    r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
    r.HandleFunc("/books/{title}", GetBook).Methods("GET")

    log.Println("http mux start..")

    http.ListenAndServe("127.0.0.1:8080", r)

    log.Println("stopping")

}

func TestMysql() {
    db, err := sql.Open("mysql", "msp:fghjiodfod2)__ABC@(112.65.210.194)/xjhh_demo?parseTime=true")

    err = db.Ping()

    createDDL := `
    CREATE TABLE users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`

    _, err = db.Exec(createDDL)
    log.Println("log println err: ", err)



    insertDML := `INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`
    result, err2 := db.Exec(insertDML, "zs", "111111", time.Now())
    userID, err3 := result.LastInsertId()
    log.Println(userID, err2, err3)

    var (
        id        int
        username  string
        password  string
        createdAt time.Time
    )

    query := `select * from users where id > ?`
    err4 := db.QueryRow(query, 0).Scan(&id, &username, &password, &createdAt)
    log.Println(err4)
    log.Println(username, password, createdAt, id)


    type user struct {
        id        int
        username  string
        password  string
        createdAt time.Time
    }
    query5 := `select * from users where id > ?`
    rows, err5 := db.Query(query5, 0)
    defer rows.Close()

    var users []user
    for rows.Next() {
        var u user
        err6 := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
        users = append(users, u)
        log.Println(err6, u)
    }
    log.Println(err5, users)
}