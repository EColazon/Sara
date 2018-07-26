package handleMySql

import (
	"fmt"
	"bytes"
	alarmMethod "handleAlarmUpload"
)

//03
type tenKeysGeter interface {
	HandleDBTenKeysGetSingle()		[]interface{}
	HandleDBTenKeysGetManny()		[]interface{}
}

//定义TenKeys数据库表结构体
//单灯阶段调光&特殊策略&节假日策略
type BuffTenKeys struct{
	uid,num int
	t01		int
	t02		int
	t03		int
	t04		int
	t05		int
	t06		int
	t07		int
	t08		int
	t09		int
	t10		int

}

const (
	sqlInHeadTenKeys = "INSERT "
	SqlUpHeadTenKeys = "UPDATE "
	sqlInTailTenKeys = " SET num=?,t01=?,t02=?,t03=?,t04=?,t05=?,t06=?,t07=?,t08=?,t09=?,t10=?,update_time=CURRENT_TIMESTAMP()"
	sqlUpTailTenKeys = " SET num=?,t01=?,t02=?,t03=?,t04=?,t05=?,t06=?,t07=?,t08=?,t09=?,t10=?,update_time=CURRENT_TIMESTAMP() WHERE num = ?"
	sqlTenKeysDltHead = "DELETE FROM "
	sqlTenKeysDltTail = " WHERE num = ?"
	sqlTenKeysSlctHead = "SELECT uid,num,t01,t02,t03,t04,t05,t06,t07,t08,t09,t10 from "
	sqlTenKeysSlctSinTail = " where num = ?"
	sqlTenKeysSlctMnyTail = " where num >= ? and num <= ?"
)

//更新数据
func HandleDBTenKeysInsert(num int, tenDatas []int, dbname string) (bool){
	var buff BuffTenKeys
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
	sqlIn.WriteString(sqlInHeadTenKeys)
	sqlIn.WriteString(dbname)
	sqlIn.WriteString(sqlInTailTenKeys)
	//拼组Update
	sqlUp.WriteString(SqlUpHeadTenKeys)
	sqlUp.WriteString(dbname)
	sqlUp.WriteString(sqlUpTailTenKeys)

	sqlElecIn := sqlIn.String()
	sqlElecUp := sqlUp.String() 

	fmt.Println("---> Insert SqlString In", sqlElecIn)
	fmt.Println("---> Insert SqlString Up", sqlElecUp)
	// sqlElecIn := "INSERT dbelec SET num=?,current=?,volt=?,pf=?,power=?,energy=?"
	// sqlElecUp := "UPDATE dbelec SET num=?,current=?,volt=?,pf=?,power=?,energy=? WHERE num = ?"
	ok := buff.HandleDBTenKeysGetSingle(num, dbname)
	if len(ok) > 0{ //数据存在->更新
		stmt, err := tx.Prepare(sqlElecUp)
		fmt.Println("---> Prepare Up")
		if err != nil{
			fmt.Println("---> Prepare fail", err)
			return false
		}
		//将参数传递到sql语句中并且执行
		res, err := stmt.Exec(num,tenDatas[0],tenDatas[1], tenDatas[2], tenDatas[3], tenDatas[4],tenDatas[5],tenDatas[6], tenDatas[7], tenDatas[8], tenDatas[9], num)
		if err != nil{
			fmt.Println("---> Exec fail", err)
			tx.Rollback()
			AlarmBuffDB[8] = 0xD4
			alarmMethod.HandleAlarmBuffParsing(AlarmBuffDB)
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
		res, err := stmt.Exec(num,tenDatas[0],tenDatas[1], tenDatas[2], tenDatas[3], tenDatas[4],tenDatas[5],tenDatas[6], tenDatas[7], tenDatas[8], tenDatas[9])
		if err != nil{
			fmt.Println("---> Exec fail", err)
			tx.Rollback()
			AlarmBuffDB[8] = 0xD4
			alarmMethod.HandleAlarmBuffParsing(AlarmBuffDB)
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
func HandleDBTenKeysDelete(num int, dbname string) (bool) {
	//准备sql语句
	sqlDelete := bytes.Buffer{}
	//拼组Insert
	sqlDelete.WriteString(sqlTenKeysDltHead)
	sqlDelete.WriteString(dbname)
	sqlDelete.WriteString(sqlTenKeysDltTail)
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
		AlarmBuffDB[8] = 0xD4
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
func (buff BuffTenKeys)HandleDBTenKeysGetSingle(num int, dbname string) ([]interface{}) {
	//准备sql语句
	sqlSelect := bytes.Buffer{}
	//拼组Insert
	sqlSelect.WriteString(sqlTenKeysSlctHead)
	sqlSelect.WriteString(dbname)
	sqlSelect.WriteString(sqlTenKeysSlctSinTail)
	sqlSlct := sqlSelect.String()

	// var buff BuffTenKeys
	var buffs []interface{}
	// var num int
	// var current,volt,pf,power,energy float64
	//执行查询语句
	// sqlSlct := "SELECT uid,t01,t02,t03,t04,t05,t06,t07,t08,t09,t10 from dbelec where num = ?"
    rows, err := DB.Query(sqlSlct, num)
    if err != nil{
		AlarmBuffDB[8] = 0xD4
		alarmMethod.HandleAlarmBuffParsing(AlarmBuffDB)
		fmt.Println("---> Select Error.")
		return buffs
	}
	
	//循环读取结果
    for rows.Next(){
        //将每一行的结果都赋值到一个user对象中
        err := rows.Scan(&buff.uid, &buff.num, &buff.t01, &buff.t02, &buff.t03, &buff.t04, &buff.t05,&buff.t06, &buff.t07, &buff.t08, &buff.t09, &buff.t10)
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
func (buff BuffTenKeys)HandleDBTenKeysGetManny(index, end int, dbname string) ([]interface{}) {
	//准备sql语句
	sqlSelect := bytes.Buffer{}
	//拼组Insert
	sqlSelect.WriteString(sqlTenKeysSlctHead)
	sqlSelect.WriteString(dbname)
	sqlSelect.WriteString(sqlTenKeysSlctMnyTail)
	sqlSlct := sqlSelect.String()

	// var buff BuffTenKeys
	// var num int
	// var current,volt,pf,power,energy float64
	//执行查询语句
	// sqlSlct := "SELECT uid,t01,t02,t03,t04,t05,t06,t07,t08,t09,t10 from dbelec where num >= ? and num <= ?"
    rows, err := DB.Query(sqlSlct, index, end)
    if err != nil{
		AlarmBuffDB[8] = 0xD4
		alarmMethod.HandleAlarmBuffParsing(AlarmBuffDB)
        fmt.Println("---> Select Error.")    
	}
	var buffs []interface{}
	//循环读取结果
    for rows.Next(){
        //将每一行的结果都赋值到一个user对象中
        err := rows.Scan(&buff.uid, &buff.num, &buff.t01, &buff.t02, &buff.t03, &buff.t04, &buff.t05,&buff.t06, &buff.t07, &buff.t08, &buff.t09, &buff.t10)
        if err != nil {
            fmt.Println("---> rows fail")
        }
        //将user追加到users的这个数组中
        buffs = append(buffs, buff)
	}
	fmt.Println("---> buffs: ", buffs)
    return buffs
}