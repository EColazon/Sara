package handleMySql

/*
说明:
	dblogcmderr			 命令解析异常记录
	dblogcmdok00		 Server->RTU	
	dblogcmdok01		 RTU->Router
	dblogcmdok02		 Router->RTU
	dblogcmdok03		 RTU->Server
	dblogcmdok04		 RTUDect
*/
import (
	"fmt"
	"bytes"
	alarmMethod "handleAlarmUpload"
)

type logGeter interface {
	HandleDBLogGetSingle()		[]interface{}
	HandleDBLogGetManny()		[]interface{}
}
//定义电参量数据库表结构体
type BuffLog struct{
	uid,flag int
	content string
	logTime string
}

const (
	sqlInHeadLog = "INSERT "
	SqlUpHeadLog = "UPDATE "
	SqlDelHeadLog = "DELETE FROM "
	sqlInTailLog = " SET flag=?,content=?,update_time=CURRENT_TIMESTAMP()"
	SqlDelTailLog = " WHERE flag = ?"
)


//插入数据
// func HandleDBLogInsert(num int, elecDatas []float64, dbname string) (bool){
func HandleDBLogInsert(flag int, content string, dbname string) (bool){
	//开启事务
	tx, err := DB.Begin()
	if err != nil{
		fmt.Println("---> tx fail")
		return false
	}
	//准备sql语句
	sqlIn := bytes.Buffer{}
	//拼组Insert
	sqlIn.WriteString(sqlInHeadLog)
	sqlIn.WriteString(dbname)
	sqlIn.WriteString(sqlInTailLog)

	sqlElecIn := sqlIn.String()

	fmt.Println("---> Insert SqlString In", sqlElecIn)
	// sqlElecIn := "INSERT dbelec SET num=?,current=?,volt=?,pf=?,power=?,energy=?"
	
 	//数据不存在->插入
	stmt, err := tx.Prepare(sqlElecIn)
	fmt.Println("---> Prepare In")
	if err != nil{
		fmt.Println("---> Prepare fail", err)
		return false
	}
	res, err := stmt.Exec(flag,content)
	if err != nil{
		fmt.Println("---> Exec fail", err)
		tx.Rollback()
		AlarmBuffDB[8] = 0xD0
		alarmMethod.HandleAlarmBuffParsing(AlarmBuffDB)
		return false
	}
	//将事务提交
	tx.Commit()
	//获得上一个插入自增的id
	fmt.Println(res.LastInsertId())
	return true
}

// 删除数据
func HandleDBLogDelete(num int, dbname string) (bool) {
    //开启事务
    tx, err := DB.Begin()
    if err != nil{
        fmt.Println("---> tx fail")
	}
	//准备sql语句
	sqlDel := bytes.Buffer{}
	//拼组Insert
	sqlDel.WriteString(SqlDelHeadLog)
	sqlDel.WriteString(dbname)
	sqlDel.WriteString(SqlDelTailLog)

	sqlDelete := sqlDel.String()
    stmt, err := tx.Prepare(sqlDelete)
    if err != nil{
        fmt.Println("---> Prepare fail")
        return false
    }
    //设置参数以及执行sql语句
    res, err := stmt.Exec(num)
    if err != nil{
		fmt.Println("---> Exec fail")
		tx.Rollback()
		AlarmBuffDB[8] = 0xD0
		alarmMethod.HandleAlarmBuffParsing(AlarmBuffDB)
        return false
    }
    //提交事务
    tx.Commit()
    //获得上一个insert的id
    fmt.Println(res.LastInsertId())
    return true
}

// 获取单条数据
func (buff BuffLog)HandleDBLogGetSingle(flag int) ([]interface{}) {
	// var buff Buff
	var buffs []interface{}
	// var num int
	// var current,volt,pf,power,energy float64
	//执行查询语句
	sqlSelect := "SELECT uid,flag,content,update_time from dblogcmderr where flag = ?"
    rows, err := DB.Query(sqlSelect, flag)
    if err != nil{
		AlarmBuffDB[8] = 0xD0
		alarmMethod.HandleAlarmBuffParsing(AlarmBuffDB)
		fmt.Println("---> Select Error.")
		return buffs
	}
	
	//循环读取结果
    for rows.Next(){
        //将每一行的结果都赋值到一个user对象中
        err := rows.Scan(&buff.uid, &buff.flag ,&buff.content, &buff.logTime)
        if err != nil {
            fmt.Println("---> rows fail")
        }
        //将user追加到users的这个数组中
        buffs = append(buffs, buff)
	}
	fmt.Println("---> buffs: ", buffs)
    return buffs
}

// 获取多条数据
func (buff BuffLog)HandleDBLogGetManny(index, end int) ([]interface{}) {
	// var buff Buff
	// var num int
	// var current,volt,pf,power,energy float64
	//执行查询语句
	sqlSelect := "SELECT uid,flag,content,update_time from dblogcmderr where uid >= ? and uid <= ?"
    rows, err := DB.Query(sqlSelect, index, end)
    if err != nil{
		AlarmBuffDB[8] = 0xD0
		alarmMethod.HandleAlarmBuffParsing(AlarmBuffDB)
        fmt.Println("---> Select Error.")    
	}
	var buffs []interface{}
	//循环读取结果
    for rows.Next(){
        //将每一行的结果都赋值到一个user对象中
        err := rows.Scan(&buff.uid, &buff.flag ,&buff.content, &buff.logTime)
        if err != nil {
            fmt.Println("---> rows fail")
        }
        //将user追加到users的这个数组中
        buffs = append(buffs, buff)
	}
	fmt.Println("---> buffs: ", buffs)
    return buffs
}