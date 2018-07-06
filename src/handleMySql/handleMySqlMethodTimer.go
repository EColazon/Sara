package handleMySql

import (
	"fmt"
	"bytes"
)

//定义电参量数据库表结构体
type BuffTimer struct{
	uid,num 	int
	lampMOpen	int
	lampMClose	int
	lampFOpen	int
	lampFClose 	int
}

const (
	sqlInHeadTimer = "INSERT "
	SqlUpHeadTimer = "UPDATE "
	sqlInTailTimer = " SET num=?,lampMOpen=?,lampMClose=?,lampFOpen=?,lampFClose=?,update_time=CURRENT_TIMESTAMP()"
	sqlUpTailTimer = " SET num=?,lampMOpen=?,lampMClose=?,lampFOpen=?,lampFClose=?,update_time=CURRENT_TIMESTAMP() WHERE num = ?"
)

//更新数据
func HandleDBSlTimerSwitchInsert(num int, timerDatas []int, dbname string) (bool){
	//开启事务
	tx, err := DB.Begin()
	if err != nil{
		fmt.Println("---> tx fail")
		return false
	}
	//准备sql语句
	sqlIn := bytes.Buffer{}
	sqlUp := bytes.Buffer{}
	//拼组Insert
	sqlIn.WriteString(sqlInHeadTimer)
	sqlIn.WriteString(dbname)
	sqlIn.WriteString(sqlInTailTimer)
	//拼组Update
	sqlUp.WriteString(SqlUpHeadTimer)
	sqlUp.WriteString(dbname)
	sqlUp.WriteString(sqlUpTailTimer)

	sqlElecIn := sqlIn.String()
	sqlElecUp := sqlUp.String() 

	fmt.Println("---> Insert SqlString In", sqlElecIn)
	fmt.Println("---> Insert SqlString Up", sqlElecUp)
	// sqlElecIn := "INSERT dbelec SET num=?,current=?,volt=?,pf=?,power=?,energy=?"
	// sqlElecUp := "UPDATE dbelec SET num=?,current=?,volt=?,pf=?,power=?,energy=? WHERE num = ?"
	ok := HandleDBSlTimerSwitchGetSingle(num)
	if len(ok) > 0{ //数据存在->更新
		stmt, err := tx.Prepare(sqlElecUp)
		fmt.Println("---> Prepare Up")
		if err != nil{
			fmt.Println("---> Prepare fail", err)
			return false
		}
		//将参数传递到sql语句中并且执行
		res, err := stmt.Exec(num,timerDatas[0],timerDatas[1], timerDatas[2], timerDatas[3], num)
		if err != nil{
			fmt.Println("---> Exec fail", err)
			tx.Rollback()
			return false
		}
		//将事务提交
		tx.Commit()
		//获得上一个插入自增的id
		fmt.Println(res.LastInsertId())
		return true
	} else { //数据不存在->插入
		stmt, err := tx.Prepare(sqlElecIn)
		fmt.Println("---> Prepare In")
		if err != nil{
			fmt.Println("---> Prepare fail", err)
			return false
		}
		res, err := stmt.Exec(num,timerDatas[0],timerDatas[1], timerDatas[2], timerDatas[3])
		if err != nil{
			fmt.Println("---> Exec fail", err)
			tx.Rollback()
			return false
		}
		//将事务提交
		tx.Commit()
		//获得上一个插入自增的id
		fmt.Println(res.LastInsertId())
		return true
	}
}

// 删除数据
func HandleDBSlTimerSwitchDelete(num int) (bool) {
    //开启事务
    tx, err := DB.Begin()
    if err != nil{
        fmt.Println("---> tx fail")
    }
	//准备sql语句
	sqlDelete := "DELETE FROM dbsltimer WHERE num = ?"
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
        return false
    }
    //提交事务
    tx.Commit()
    //获得上一个insert的id
    fmt.Println(res.LastInsertId())
    return true
}

// 获取单条数据
func HandleDBSlTimerSwitchGetSingle(num int) ([]interface{}) {
	var buff BuffTimer
	var buffs []interface{}
	// var num int
	// var current,volt,pf,power,energy float64
	//执行查询语句
	sqlSelect := "SELECT uid,num,lampMOpen,lampMClose,lampFOpen,lampFClose from dbsltimer where num = ?"
    rows, err := DB.Query(sqlSelect, num)
    if err != nil{
		fmt.Println("---> Select Error.")
		return buffs
	}
	
	//循环读取结果
    for rows.Next(){
        //将每一行的结果都赋值到一个user对象中
        err := rows.Scan(&buff.uid, &buff.num, &buff.lampMOpen, &buff.lampMClose, &buff.lampFOpen, &buff.lampFClose)
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
func HandleDBSlTimerSwitchGetManny(index, end int) ([]interface{}) {
	var buff BuffTimer
	// var num int
	// var current,volt,pf,power,energy float64
	//执行查询语句
	sqlSelect := "SELECT uid,num,lampMOpen,lampMClose,lampFOpen,lampFClose from dbsltimer where num >= ? and num <= ?"
    rows, err := DB.Query(sqlSelect, index, end)
    if err != nil{
        fmt.Println("---> Select Error.")    
	}
	var buffs []interface{}
	//循环读取结果
    for rows.Next(){
        //将每一行的结果都赋值到一个user对象中
        err := rows.Scan(&buff.uid, &buff.num, &buff.lampMOpen, &buff.lampMClose, &buff.lampFOpen, &buff.lampFClose)
        if err != nil {
            fmt.Println("---> rows fail")
        }
        //将user追加到users的这个数组中
        buffs = append(buffs, buff)
	}
	fmt.Println("---> buffs: ", buffs)
    return buffs
}