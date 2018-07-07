package handleMySql

import (
	"fmt"
	"bytes"
)

type elecGeter interface {
	HandleDBElecGetSingle()		[]interface{}
	HandleDBElecGetManny()		[]interface{}
}
//定义电参量数据库表结构体
type Buff struct{
	uid,num int
	current,volt,pf,power,energy float64
}

const (
	sqlInHead = "INSERT "
	SqlUpHead = "UPDATE "
	sqlInTail = " SET num=?,current=?,volt=?,pf=?,power=?,energy=?,update_time=CURRENT_TIMESTAMP()"
	sqlUpTail = " SET num=?,current=?,volt=?,pf=?,power=?,energy=?,update_time=CURRENT_TIMESTAMP() WHERE num = ?"
)

//创建表
func HandleDBCreateTable() {
	
	//创建电参量表
	sqlDBElec := "CREATE TABLE IF NOT EXISTS dbelec(" +
		"uid int AUTO_INCREMENT NOT NULL ," +
		"num int NOT NULL ," +
		"current float NOT NULL ," +
		"volt float NOT NULL ," +
		"pf float NOT NULL ," +
		"power float NOT NULL ," +
		"energy float NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
		") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBElecdrop := "DROP TABLE dbelec"
	DB.Exec(sqlDBElecdrop)	
	fmt.Println("---> sql: 01 ")
	DB.Exec(sqlDBElec)

	//创建节点表
	sqlDBNode := "CREATE TABLE IF NOT EXISTS dbnode(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"num INT NOT NULL ," +
		"stateNode INT NOT NULL ," +
		"timeTrigeH INT NOT NULL ," +
		"timeTrighM INT NOT NULL ," +
		"stateLight INT NOT NULL ," +
		"timeRecoverH INT NOT NULL ," +
		"timeRecoverM INT NOT NULL ," +
		"timeOpenH INT NOT NULL ," +
		"timeOpenM INT NOT NULL ," +
		"timeCloseH INT NOT NULL ," +
		"timeCloseM INT NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBNodedrop := "DROP TABLE dbinode"
	DB.Exec(sqlDBNodedrop)	
	DB.Exec(sqlDBNode)
	fmt.Println("---> sql: 02 ")
	
	//连续三个相同
	//创建单灯阶段调光表
	sqlDBStagePwm := "CREATE TABLE IF NOT EXISTS dbstagepwm(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"num INT NOT NULL ," +
		"t01 INT NOT NULL ," +
		"t02 INT NOT NULL ," +
		"t03 INT NOT NULL ," +
		"t04 INT NOT NULL ," +
		"t05 INT NOT NULL ," +
		"t06 INT NOT NULL ," +
		"t07 INT NOT NULL ," +
		"t08 INT NOT NULL ," +
		"t09 INT NOT NULL ," +
		"t10 INT NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBStagePwmdrop := "DROP TABLE dbstagepwm"
	DB.Exec(sqlDBStagePwmdrop)	
	fmt.Println("---> sql: 03 ")
	DB.Exec(sqlDBStagePwm)
	//创建特殊时间策略表
	sqlDBSpecial := "CREATE TABLE IF NOT EXISTS dbspecial(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"num INT NOT NULL ," +
		"t01 INT NOT NULL ," +
		"t02 INT NOT NULL ," +
		"t03 INT NOT NULL ," +
		"t04 INT NOT NULL ," +
		"t05 INT NOT NULL ," +
		"t06 INT NOT NULL ," +
		"t07 INT NOT NULL ," +
		"t08 INT NOT NULL ," +
		"t09 INT NOT NULL ," +
		"t10 INT NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBSpecialdrop := "DROP TABLE dbspecial"
	DB.Exec(sqlDBSpecialdrop)	
	fmt.Println("---> sql: 04 ")
	DB.Exec(sqlDBSpecial)
	//创建节假日时间策略表
	sqlDBHoliday := "CREATE TABLE IF NOT EXISTS dbholiday(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"num INT NOT NULL ," +
		"t01 INT NOT NULL ," +
		"t02 INT NOT NULL ," +
		"t03 INT NOT NULL ," +
		"t04 INT NOT NULL ," +
		"t05 INT NOT NULL ," +
		"t06 INT NOT NULL ," +
		"t07 INT NOT NULL ," +
		"t08 INT NOT NULL ," +
		"t09 INT NOT NULL ," +
		"t10 INT NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBHolidaydrop := "DROP TABLE dbholiday"
	DB.Exec(sqlDBHolidaydrop)	
	fmt.Println("---> sql: 05 ")
	DB.Exec(sqlDBHoliday)
	
	//创建单灯pwm表
	sqlDBSlPwm := "CREATE TABLE IF NOT EXISTS dbslpwm(" +
		"uid int AUTO_INCREMENT NOT NULL ," +
		"num int NOT NULL ," +
		"t01 float NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
		") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBSlPwmdrop := "DROP TABLE dbslpwm"
	DB.Exec(sqlDBSlPwmdrop)	
	fmt.Println("---> sql: 06 ")
	DB.Exec(sqlDBSlPwm)

	//创建单灯定时开关时间表
	sqlDBSlTimer := "CREATE TABLE IF NOT EXISTS dbsltimer(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"num INT NOT NULL ," +
		"lampMOpen INT NOT NULL ," +
		"lampMClose INT NOT NULL ," +
		"lampFOpen INT NOT NULL ," +
		"lampFClose INT NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBSlTimerdrop := "DROP TABLE dbsltimer"
	DB.Exec(sqlDBSlTimerdrop)	
	fmt.Println("---> sql: 07 ")
	DB.Exec(sqlDBSlTimer)
	
	// 连续四个相同结构表
	//创建单灯组类型表
	sqlDBSlGType := "CREATE TABLE IF NOT EXISTS dbslgtype(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"num INT NOT NULL ," +
		"lampGType INT NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBSlGTypeDrop := "DROP TABLE dbslgtype"
	DB.Exec(sqlDBSlGTypeDrop)	
	fmt.Println("---> sql: 08 ")
	DB.Exec(sqlDBSlGType)
	//创建单灯组号表
	sqlDBSlGNum := "CREATE TABLE IF NOT EXISTS dbslgnum(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"num INT NOT NULL ," +
		"lampGNum INT NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBSlGNumDrop := "DROP TABLE dbslgnum"
	DB.Exec(sqlDBSlGNumDrop)	
	fmt.Println("---> sql: 09 ")
	DB.Exec(sqlDBSlGNum)
	//创建单灯继电器状态表
	sqlDBSlRelayST := "CREATE TABLE IF NOT EXISTS dbslrelayst(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"num INT NOT NULL ," +
		"lampRelayST INT NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBSlRelaySTDrop := "DROP TABLE dbslrelayst"
	DB.Exec(sqlDBSlRelaySTDrop)	
	fmt.Println("---> sql: 10 ")
	DB.Exec(sqlDBSlRelayST)
	//创建单灯手动开关状态表
	sqlDBSlHandST := "CREATE TABLE IF NOT EXISTS dbslhandst(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"num INT NOT NULL ," +
		"lampHandST INT NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBSlHandSTDrop := "DROP TABLE dbslhandst"
	DB.Exec(sqlDBSlHandSTDrop)	
	fmt.Println("---> sql: 11 ")
	DB.Exec(sqlDBSlHandST)



}

//更新数据
func HandleDBElecInsert(num int, elecDatas []float64, dbname string) (bool){
	var buff Buff
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
	ok := buff.HandleDBElecGetSingle(num)
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
func (buff Buff)HandleDBElecGetSingle(num int) ([]interface{}) {
	// var buff Buff
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
func (buff Buff)HandleDBElecGetManny(index, end int) ([]interface{}) {
	// var buff Buff
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