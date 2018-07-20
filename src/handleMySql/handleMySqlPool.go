package handleMySql

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "net/http"
	
)

//var DB *sql.DB
 
func InitDB() {
    DB, _ = sql.Open("mysql", "root:hello520Sara@tcp(127.0.0.1:3306)/dbtest?charset=utf8")
    DB.SetMaxOpenConns(2000)
    DB.SetMaxIdleConns(1000)
    DB.Ping()
}
 
func HandleMySqlPool() {
    startHttpServer()
}
 
func startHttpServer() {
    http.HandleFunc("/pool", pool)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
 
func pool(w http.ResponseWriter, r *http.Request) {
    rows, err := DB.Query("SELECT * FROM user")
    defer rows.Close()
    checkErr(err)
 
    columns, _ := rows.Columns()
    scanArgs := make([]interface{}, len(columns))
    values := make([]interface{}, len(columns))
    for j := range values {
        scanArgs[j] = &values[j]
    }
 
    record := make(map[string]string)
    for rows.Next() {
        //将行数据保存到record字典
        err = rows.Scan(scanArgs...)
        for i, col := range values {
            if col != nil {
                record[columns[i]] = string(col.([]byte))
            }
        }
    }
 
    fmt.Println(record)
    fmt.Fprintln(w, "finish")
}
 
func checkErr(err error) {
    if err != nil {
        fmt.Println(err)
        panic(err)
    }
}