package handleMySql

import (
	"fmt"
	"bytes"
)

//03

//定义OneKeys数据库表结构体
// 单灯PWM数据表&其他&单灯组类型&单灯组号&单灯继电器状态&单灯手动开关状态
type BuffOneKeys struct{
	uid,num int
	t01		int

}

const (
	sqlInHeadOneKeys = "INSERT "
	SqlUpHeadOneKeys = "UPDATE "
	sqlInTailOneKeys = " SET num=?,t01=?,update_time=CURRENT_TIMESTAMP()"
	sqlUpTailOneKeys = " SET num=?,t01=?,update_time=CURRENT_TIMESTAMP() WHERE num = ?"
	sqlOneKeysDltHead = "DELETE FROM "
	sqlOneKeysDltTail = " WHERE num = ?"
	sqlOneKeysSlctHead = "SELECT uid,num,t01 from "
	sqlOneKeysSlctSinTail = " where num = ?"
	sqlOneKeysSlctMnyTail = " where num >= ? and num <= ?"
)

//更新数据
func HandleDBOneKeysInsert(num int, oneDatas []int, dbname string) (bool){
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
	sqlIn.WriteString(sqlInHeadOneKeys)
	sqlIn.WriteString(dbname)
	sqlIn.WriteString(sqlInTailOneKeys)
	//拼组Update
	sqlUp.WriteString(SqlUpHeadOneKeys)
	sqlUp.WriteString(dbname)
	sqlUp.WriteString(sqlUpTailOneKeys)

	sqlElecIn := sqlIn.String()
	sqlElecUp := sqlUp.String() 

	fmt.Println("---> Insert SqlString In", sqlElecIn)
	fmt.Println("---> Insert SqlString Up", sqlElecUp)
	// sqlElecIn := "INSERT dbelec SET num=?,current=?,volt=?,pf=?,power=?,energy=?"
	// sqlElecUp := "UPDATE dbelec SET num=?,current=?,volt=?,pf=?,power=?,energy=? WHERE num = ?"
	ok := HandleDBOneKeysGetSingle(num, dbname)
	if len(ok) > 0{ //数据存在->更新
		stmt, err := tx.Prepare(sqlElecUp)
		fmt.Println("---> Prepare Up")
		if err != nil{
			fmt.Println("---> Prepare fail", err)
			return false
		}
		//将参数传递到sql语句中并且执行
		res, err := stmt.Exec(num,oneDatas[0], num)
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
		res, err := stmt.Exec(num,oneDatas[0])
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
func HandleDBOneKeysDelete(num int, dbname string) (bool) {
	//准备sql语句
	sqlDelete := bytes.Buffer{}
	//拼组Insert
	sqlDelete.WriteString(sqlOneKeysDltHead)
	sqlDelete.WriteString(dbname)
	sqlDelete.WriteString(sqlOneKeysDltTail)
	sqlDlt := sqlDelete.String()
    //开启事务
    tx, err := DB.Begin()
    if err != nil{
        fmt.Println("---> tx fail")
    }
	//准备sql语句
	// sqlDlt := "DELETE FROM dbelec WHERE num = ?"
    stmt, err := tx.Prepare(sqlDlt)
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
func HandleDBOneKeysGetSingle(num int, dbname string) ([]interface{}) {
	//准备sql语句
	sqlSelect := bytes.Buffer{}
	//拼组Insert
	sqlSelect.WriteString(sqlOneKeysSlctHead)
	sqlSelect.WriteString(dbname)
	sqlSelect.WriteString(sqlOneKeysSlctSinTail)
	sqlSlct := sqlSelect.String()

	var buff BuffOneKeys
	var buffs []interface{}
	// var num int
	// var current,volt,pf,power,energy float64
	//执行查询语句
	// sqlSlct := "SELECT uid,t01,t02,t03,t04,t05,t06,t07,t08,t09,t10 from dbelec where num = ?"
    rows, err := DB.Query(sqlSlct, num)
    if err != nil{
		fmt.Println("---> Select Error.")
		return buffs
	}
	
	//循环读取结果
    for rows.Next(){
        //将每一行的结果都赋值到一个user对象中
        err := rows.Scan(&buff.uid, &buff.num, &buff.t01)
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
func HandleDBOneKeysGetManny(index, end int, dbname string) ([]interface{}) {
	//准备sql语句
	sqlSelect := bytes.Buffer{}
	//拼组Insert
	sqlSelect.WriteString(sqlOneKeysSlctHead)
	sqlSelect.WriteString(dbname)
	sqlSelect.WriteString(sqlOneKeysSlctMnyTail)
	sqlSlct := sqlSelect.String()

	var buff BuffOneKeys
	// var num int
	// var current,volt,pf,power,energy float64
	//执行查询语句
	// sqlSlct := "SELECT uid,t01,t02,t03,t04,t05,t06,t07,t08,t09,t10 from dbelec where num >= ? and num <= ?"
    rows, err := DB.Query(sqlSlct, index, end)
    if err != nil{
        fmt.Println("---> Select Error.")    
	}
	var buffs []interface{}
	//循环读取结果
    for rows.Next(){
        //将每一行的结果都赋值到一个user对象中
        err := rows.Scan(&buff.uid, &buff.num, &buff.t01)
        if err != nil {
            fmt.Println("---> rows fail")
        }
        //将user追加到users的这个数组中
        buffs = append(buffs, buff)
	}
	fmt.Println("---> buffs: ", buffs)
    return buffs
}