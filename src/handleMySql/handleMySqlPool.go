package handleMySql

//数据库连接池测试
 
import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "net/http"
    "strings"
	
)

//数据库配置
const (
    userName = "root"
    password = "hello520Sara"
    ip = "127.0.0.1"
    port = "3306"
    dbName = "dbtest"
)
//Db数据库连接池
var DB *sql.DB

//注意方法名大写，就是public
func init()  {
    //构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
    path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

    //打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
    DB, _ = sql.Open("mysql", path)
    //设置数据库最大连接数
    DB.SetConnMaxLifetime(100)
    //设置上数据库最大闲置连接数
    DB.SetMaxIdleConns(10)
    //验证连接
    if err := DB.Ping(); err != nil{
        fmt.Println("opon database fail")
        return
    }
    fmt.Println("connnect success")
}

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