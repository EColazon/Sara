package handleMySql

import (
	"fmt"
	"bytes"
)

type lampGeter interface {
	HandleDBSlLampGetSingle()		[]interface{}
	HandleDBSlLampGetManny()		[]interface{}
}
//定义电参量数据库表结构体
type BuffLamp struct{
	uid 	int
	lampNum int
	lampNumGroup int
	lAdvV int
	lAdvI int
	lAdvP int
	lAdvPF int
	lAuxV int
	lAuxI int
	lAuxP int
	lAuxPF int
	lAdvPwm int
	lAuxPwm int
	lStateRelayBT int
	lModeTX int
	lModeRX int
	lTimeAlarm int
	lRelayChange int
	lFlagEX int
	lStateAlarm int
	lBNetAddr int
	lAdvPower int
	lAuxPower int
	lampHigherV int
	lampLowerV int
	lampHigherI int
	lampLowerI int
	lampHigherP int
	lampLowerP int
	lampHigherPF int
	lampLowerPF int
	lFlagSetNum int
	lFlagSetAdu int
	lChecksum int
}

const (
	sqlInHeadLamp = "INSERT "
	SqlUpHeadLamp = "UPDATE "
	sqlInTailLamp = " SET lampNum=?,lampNumGroup=?,lAdvV=?,lAdvI=?,lAdvP=?,lAdvPF=?,lAuxV=?,lAuxI=?,lAuxP=?,lAuxPF=?,lAdvPwm=?,lAuxPwm=?,lStateRelayBT=?,lModeTX=?,lModeRX=?,lTimeAlarm=?,lRelayChange=?,lFlagEX=?,lStateAlarm=?,lBNetAddr=?,lAdvPower=?,lAuxPower=?,lampHigherV=?,lampLowerV=?,lampHigherI=?,lampLowerI=?,lampHigherP=?,lampLowerP=?,lampHigherPF=?,lampLowerPF=?,lFlagSetNum=?,lFlagSetAdu=?,lChecksum=?,update_time=CURRENT_TIMESTAMP()"
	sqlUpTailLamp = " SET lampNum=?,lampNumGroup=?,lAdvV=?,lAdvI=?,lAdvP=?,lAdvPF=?,lAuxV=?,lAuxI=?,lAuxP=?,lAuxPF=?,lAdvPwm=?,lAuxPwm=?,lStateRelayBT=?,lModeTX=?,lModeRX=?,lTimeAlarm=?,lRelayChange=?,lFlagEX=?,lStateAlarm=?,lBNetAddr=?,lAdvPower=?,lAuxPower=?,lampHigherV=?,lampLowerV=?,lampHigherI=?,lampLowerI=?,lampHigherP=?,lampLowerP=?,lampHigherPF=?,lampLowerPF=?,lFlagSetNum=?,lFlagSetAdu=?,lChecksum=?,update_time=CURRENT_TIMESTAMP() WHERE num = ?"
)

//更新数据
func HandleDBSlLampInsert(num int, timerDatas []int, dbname string) (bool){
	var buff BuffLamp
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
	ok := buff.HandleDBSlLampGetSingle(num)
	if len(ok) > 0{ //数据存在->更新
		stmt, err := tx.Prepare(sqlElecUp)
		fmt.Println("---> Prepare Up")
		if err != nil{
			fmt.Println("---> Prepare fail", err)
			return false
		}
		//将参数传递到sql语句中并且执行
		res, err := stmt.Exec(num,timerDatas[0],timerDatas[1], timerDatas[2], timerDatas[3], timerDatas[4],timerDatas[5], timerDatas[6], timerDatas[7],timerDatas[8],timerDatas[9], timerDatas[10], timerDatas[11],timerDatas[12],timerDatas[13], timerDatas[14], timerDatas[15],timerDatas[16],timerDatas[17], timerDatas[18], timerDatas[19],timerDatas[20],timerDatas[21], timerDatas[22], timerDatas[23],timerDatas[24],timerDatas[25], timerDatas[26], timerDatas[27],timerDatas[28],timerDatas[29], timerDatas[30], timerDatas[31],timerDatas[32],num)
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
		res, err := stmt.Exec(num,timerDatas[0],timerDatas[1], timerDatas[2], timerDatas[3], timerDatas[4],timerDatas[5], timerDatas[6], timerDatas[7],timerDatas[8],timerDatas[9], timerDatas[10], timerDatas[11],timerDatas[12],timerDatas[13], timerDatas[14], timerDatas[15],timerDatas[16],timerDatas[17], timerDatas[18], timerDatas[19],timerDatas[20],timerDatas[21], timerDatas[22], timerDatas[23],timerDatas[24],timerDatas[25], timerDatas[26], timerDatas[27],timerDatas[28],timerDatas[29], timerDatas[30], timerDatas[31],timerDatas[32])
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
func HandleDBSlLampDelete(num int) (bool) {
    //开启事务
    tx, err := DB.Begin()
    if err != nil{
        fmt.Println("---> tx fail")
    }
	//准备sql语句
	sqlDelete := "DELETE FROM dblamp WHERE num = ?"
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
func (buff BuffLamp)HandleDBSlLampGetSingle(num int) ([]interface{}) {
	// var buff BuffTimer
	var buffs []interface{}
	// var num int
	// var current,volt,pf,power,energy float64
	//执行查询语句
	sqlSelect := "SELECT uid,lampNum,lampNumGroup,lAdvV,lAdvI,lAdvP,lAdvPF,lAuxV,lAuxI,lAuxP,lAuxPF,lAdvPwm,lAuxPwm,lStateRelayBT,lModeTX,lModeRX,lTimeAlarm,lRelayChange,lFlagEX,lStateAlarm,lBNetAddr,lAdvPower,lAuxPower,lampHigherV,lampLowerV,lampHigherI,lampLowerI,lampHigherP,lampLowerP,lampHigherPF,lampLowerPF,lFlagSetNum,lFlagSetAdu,lChecksum from dbsltimer where num = ?"
    rows, err := DB.Query(sqlSelect, num)
    if err != nil{
		fmt.Println("---> Select Error.")
		return buffs
	}
	
	//循环读取结果
    for rows.Next(){
        //将每一行的结果都赋值到一个user对象中
        err := rows.Scan(&buff.uid, &buff.lampNum,&buff.lampNumGroup,&buff.lAdvV,&buff.lAdvI,&buff.lAdvP,&buff.lAdvPF,&buff.lAuxV,&buff.lAuxI,&buff.lAuxP,&buff.lAuxPF,&buff.lAdvPwm,&buff.lAuxPwm,&buff.lStateRelayBT,&buff.lModeTX,&buff.lModeRX,&buff.lTimeAlarm,&buff.lRelayChange,&buff.lFlagEX,&buff.lStateAlarm,&buff.lBNetAddr,&buff.lAdvPower,&buff.lAuxPower,&buff.lampHigherV,&buff.lampLowerV,&buff.lampHigherI,&buff.lampLowerI,&buff.lampHigherP,&buff.lampLowerP,&buff.lampHigherPF,&buff.lampLowerPF,&buff.lFlagSetNum,&buff.lFlagSetAdu,&buff.lChecksum)
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
func (buff BuffLamp)HandleDBSlLampManny(index, end int) ([]interface{}) {
	// var buff BuffTimer
	// var num int
	// var current,volt,pf,power,energy float64
	//执行查询语句
	sqlSelect := "SELECT uid,lampNum,lampNumGroup,lAdvV,lAdvI,lAdvP,lAdvPF,lAuxV,lAuxI,lAuxP,lAuxPF,lAdvPwm,lAuxPwm,lStateRelayBT,lModeTX,lModeRX,lTimeAlarm,lRelayChange,lFlagEX,lStateAlarm,lBNetAddr,lAdvPower,lAuxPower,lampHigherV,lampLowerV,lampHigherI,lampLowerI,lampHigherP,lampLowerP,lampHigherPF,lampLowerPF,lFlagSetNum,lFlagSetAdu,lChecksum from dbsltimer where num >= ? and num <= ?"
    rows, err := DB.Query(sqlSelect, index, end)
    if err != nil{
        fmt.Println("---> Select Error.")    
	}
	var buffs []interface{}
	//循环读取结果
    for rows.Next(){
        //将每一行的结果都赋值到一个user对象中
        err := rows.Scan(&buff.uid, &buff.lampNum,&buff.lampNumGroup,&buff.lAdvV,&buff.lAdvI,&buff.lAdvP,&buff.lAdvPF,&buff.lAuxV,&buff.lAuxI,&buff.lAuxP,&buff.lAuxPF,&buff.lAdvPwm,&buff.lAuxPwm,&buff.lStateRelayBT,&buff.lModeTX,&buff.lModeRX,&buff.lTimeAlarm,&buff.lRelayChange,&buff.lFlagEX,&buff.lStateAlarm,&buff.lBNetAddr,&buff.lAdvPower,&buff.lAuxPower,&buff.lampHigherV,&buff.lampLowerV,&buff.lampHigherI,&buff.lampLowerI,&buff.lampHigherP,&buff.lampLowerP,&buff.lampHigherPF,&buff.lampLowerPF,&buff.lFlagSetNum,&buff.lFlagSetAdu,&buff.lChecksum)
        if err != nil {
            fmt.Println("---> rows fail")
        }
        //将user追加到users的这个数组中
        buffs = append(buffs, buff)
	}
	fmt.Println("---> buffs: ", buffs)
    return buffs
}