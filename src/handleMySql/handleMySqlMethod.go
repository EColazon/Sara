package handleMySql

import (
	"fmt"
	"bytes"
)

//定义电参量数据库表结构体
type Buff struct{
	uid,num int
	current,volt,pf,power,energy float64
}

const (
	sqlInHead = "INSERT "
	SqlUpHead = "UPDATE "
	sqlInTail = " SET num=?,current=?,volt=?,pf=?,power=?,energy=?"
	sqlUpTail = " SET num=?,current=?,volt=?,pf=?,power=?,energy=? WHERE num = ?"
)

//创建电参量表
func HandleDBCreateTable() {

	sql := "CREATE TABLE IF NOT EXISTS dbelec(" +
  			"uid int AUTO_INCREMENT NOT NULL ," +
  			"num int NOT NULL ," +
  			"current float NOT NULL ," +
  			"volt float NOT NULL ," +
  			"pf float NOT NULL ," +
  			"power float NOT NULL ," +
  			"energy float NOT NULL ," +
  			"update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
  			"create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
  			"PRIMARY KEY (uid)" +
			") engine=innodb DEFAULT charset=utf8mb4;"
	sqldrop := "DROP TABLE dbelec"
	DB.Exec(sqldrop)	
	fmt.Println("---> sql: ", sql)
	DB.Exec(sql)
}

//更新数据
func HandleDBElecInsert(num int, elecDatas []float64, dbname string) (bool){
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
	sqlIn.WriteString(sqlInHead)
	sqlIn.WriteString(dbname)
	sqlIn.WriteString(sqlInTail)
	//拼组Update
	sqlUp.WriteString(SqlUpHead)
	sqlUp.WriteString(dbname)
	sqlUp.WriteString(sqlUpTail)

	sqlElecIn := sqlIn.String()
	sqlElecUp := sqlUp.String() 

	fmt.Println("---> Insert SqlString In", sqlElecIn)
	fmt.Println("---> Insert SqlString Up", sqlElecUp)
	// sqlElecIn := "INSERT dbelec SET num=?,current=?,volt=?,pf=?,power=?,energy=?"
	// sqlElecUp := "UPDATE dbelec SET num=?,current=?,volt=?,pf=?,power=?,energy=? WHERE num = ?"
	ok := HandleDBElecGetSingle(num)
	if len(ok) > 0{ //数据存在->更新
		stmt, err := tx.Prepare(sqlElecUp)
		fmt.Println("---> Prepare Up")
		if err != nil{
			fmt.Println("---> Prepare fail", err)
			return false
		}
		//将参数传递到sql语句中并且执行
		res, err := stmt.Exec(num,elecDatas[0],elecDatas[1], elecDatas[2], elecDatas[3], elecDatas[4], num)
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
		res, err := stmt.Exec(num,elecDatas[0], elecDatas[1], elecDatas[2], elecDatas[3], elecDatas[4])
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
func HandleDBElecDelete(num int) (bool) {
    //开启事务
    tx, err := DB.Begin()
    if err != nil{
        fmt.Println("---> tx fail")
    }
	//准备sql语句
	sqlDelete := "DELETE FROM dbelec WHERE num = ?"
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
func HandleDBElecGetSingle(num int) ([]interface{}) {
	var buff Buff
	var buffs []interface{}
	// var num int
	// var current,volt,pf,power,energy float64
	//执行查询语句
	sqlSelect := "SELECT uid,num,current,volt,pf,power,energy from dbelec where num = ?"
    rows, err := DB.Query(sqlSelect, num)
    if err != nil{
		fmt.Println("---> Select Error.")
		return buffs
	}
	
	//循环读取结果
    for rows.Next(){
        //将每一行的结果都赋值到一个user对象中
        err := rows.Scan(&buff.uid, &buff.num, &buff.current, &buff.volt, &buff.pf, &buff.power, &buff.energy)
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
func HandleDBElecGetManny(index, end int) ([]interface{}) {
	var buff Buff
	// var num int
	// var current,volt,pf,power,energy float64
	//执行查询语句
	sqlSelect := "SELECT uid,num,current,volt,pf,power,energy from dbelec where num >= ? and num <= ?"
    rows, err := DB.Query(sqlSelect, index, end)
    if err != nil{
        fmt.Println("---> Select Error.")    
	}
	var buffs []interface{}
	//循环读取结果
    for rows.Next(){
        //将每一行的结果都赋值到一个user对象中
        err := rows.Scan(&buff.uid, &buff.num, &buff.current, &buff.volt, &buff.pf, &buff.power, &buff.energy)
        if err != nil {
            fmt.Println("---> rows fail")
        }
        //将user追加到users的这个数组中
        buffs = append(buffs, buff)
	}
	fmt.Println("---> buffs: ", buffs)
    return buffs
}