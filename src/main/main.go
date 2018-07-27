package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	// "strconv"
	// "bytes"
	// "reflect"
	// "handleCmdsManages"
	// "timefunc"
	// "gyjson"
	"time"
	/*
	"handleShared"
	"strconv"
	*/
	/*
	"chgoroutines"
	"godistrabutor"
	"gyrwio"
	"interfaces"
	
	*/
	//"goTcpServer"
	//"goredis"
	//"workerpool"
	//"handleRedis"
	"handleMySql"
	Sheard "handleShared"
	
)

func main() {
	fmt.Println("Hello Go")
	sig := make(chan os.Signal, 2)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	Sheard.ExecShell("ls")

	data := []int{0x01, 0x02, 34, 0x05, 0x06}
	
	strFinal := Sheard.Slice2String(data)
	fmt.Println("---> str: ", strFinal)

	OK := handleMySql.HandleDBLogDelete(7, "dbelec")
	fmt.Println("---> Del OK：", OK)	

	timeNowHour := time.Now().Hour()
	timeNowMinute := time.Now().Minute()
	timeNowWeek	  := time.Now().Weekday().String()
	if timeNowHour == 11 {
		fmt.Println("---> HHH.")
	}
	if timeNowWeek == "Friday" {
		fmt.Println("---> FFF.")
	}

	fmt.Println("---> TimeNow: ", timeNowHour, timeNowMinute, timeNowWeek)
	


	// sqlSelect := "SELECT uid,num,current,volt,pf,power,energy from dbelec where num = ?"
	// sqlSelect := "select * from dbelec where DATE_SUB(CURDATE(), INTERVAL  1 MONTH) <= date(update_time) "
	// rows, _ := handleMySql.DB.Query(sqlSelect)
	// fmt.Println("---> row: ", rows)
	
	//循环读取结果
    // for rows.Next(){
    //     //将每一行的结果都赋值到一个user对象中
    //     err := rows.Scan(&buff.uid, &buff.num, &buff.current, &buff.volt, &buff.pf, &buff.power, &buff.energy)
    //     if err != nil {
    //         fmt.Println("---> rows fail")
    //     }
    //     //将user追加到users的这个数组中
    //     buffs = append(buffs, buff)
	// }
	// fmt.Println("---> buffs: ", buffs)

	// for i := 1; i <= 256; i++ {
	// 	a := So.DecimalToAny(i, 16, 8)
	// 	fmt.Println("---> stra: ", a)
	// }
	// arraya := []int{1, 2, 3, 4, 56, 7}
	// arrayb := [3]int{1, 2, 3}
	// slice := make([]int, 3, 10)
	// fmt.Println("---> AAA--- > ", reflect.TypeOf(arraya))
	// fmt.Println("---> BBB--- > ", reflect.TypeOf(arrayb))
	// fmt.Println("---> SSS--- > ", reflect.TypeOf(slice))
	//命令解析模块测试
	// timefunc.TimeFunc(handleCmdsManagesSR.Cmd2FParsing)
	// timefunc.TimeFunc(handleCmdsManagesSR.Cmd33Parsing)
	// timefunc.TimeFunc(forRange)
	
	// 命令解析select-channel测试
	
	
	// fmt.Println("---> 000 ", handleCmdsManagesSR.GZigbeeNode)
	

	

	// // // czbNode = append(czbNode, zbNode)
	// // fmt.Println("---> zbNode: ", zbNode)
	// // fmt.Println("---> czbNode: ", czbNode)
	// // for i := 0; i < 5; i++ {
	// 	// zbNode.id = i
	// // }
	// // zbNode.ZigbeeNode72052SetNodeInductTrigeTimeActer(0, 1, 1, 1, 1)
		
	// var cmd  handleCmdsManagesSR.CmdChannel
	// fmt.Println("---Start 33 ---")
	// for i := 0; i < 1; i++ {
	// 	go handleCmdsManagesSR.Cmd33Parsing()
	// 	fmt.Println("---> 33 id : ", i)
	// }
	
	// // time.Sleep(1*time.Second)
	// fmt.Println("---Start 2F ---")
	// for i := 0; i < 1; i++ {
	// 	go handleCmdsManagesSR.Cmd2FParsing()
	// 	fmt.Println("---> 2f id : ", i)
	// }
	
	
	// cmd.HandleCmdGeter()
	// time.Sleep(10*time.Second)
	// fmt.Println("----------------------------------------------> ")
	// for i := 0; i < 10; i++ {
	// 	go handleCmdsManagesSR.Cmd33Parsing()
	// 	fmt.Println("---> 33 id : ", i)
	// }
	
	// // time.Sleep(1*time.Second)
	// fmt.Println("---Start 2F ---")
	// for i := 0; i < 10; i++ {
	// 	go handleCmdsManagesSR.Cmd2FParsing()
	// 	fmt.Println("---> 2f id : ", i)
	// }

	// 超时重传测试
	/*
	var ccc handleCmdsManages.CmdZigbeeChannel
	data := []int{1, 2, 34, 5, 6}

	go handleCmdsManages.HandleCmdSendOnce(data)
	var mm = make(map[int]int, 1024)
	mm[1] = 1
	mm[2] = 2
	fmt.Println("---> mm ",len(mm), mm)
	
	ccc.HandleCmdRetransmissionDistributer()

	// for i := 0; i < 5; i++ {

	// 	handleCmdsManagesSR.HandleCmdDistributerRetransmission(0)
	// }
	
	time.Sleep(5*time.Second)
	<-sig
	// <-handleCmdsManagesSR.ChExit

	*/

	
	/*

	data := []int{0x01, 0x02, 0x03, 0x04, 0x05, 0x06}

	handleShared.HandleSharedCmdOk(22, data, 0xff)

	//文件读写测试
	tNow := time.Now()
	fmt.Println("---> tNow.Date()", tNow.String()[:19])
	timeYMD := tNow.String()[:10]
	timeHMS := tNow.String()[11:19]
	fmt.Println("---> tNow.format ",reflect.TypeOf(timeYMD), timeYMD, timeHMS)

	handleShared.RecordStrInfos("Rlog", "Hello LaoJiang")
	//调用Linux命令测试
	fmt.Println("---> ", reflect.TypeOf(`dh`))
	handleShared.ExecShell(`df -h`)

	//CRC32测试
	var crc32Buffer bytes.Buffer
	for i := 0; i < len(data); i++ {
		crc32Buffer.WriteString(strconv.Itoa(data[i]))
	}
	crc32data := crc32Buffer.String()
	fmt.Println("---> crc32DataJoin: ", crc32data)
	handleShared.ExecCRC32(crc32data)
	
	//MD5测试
	var md5Buffer bytes.Buffer
	for i := 0; i < len(data); i++ {
		md5Buffer.WriteString(strconv.Itoa(data[i]))
	}
	md5data := md5Buffer.String()
	fmt.Println("---> md5DataJoin: ", md5data)
	handleShared.ExecMD5(md5data)
	//EEPROM测试
    fmt.Println("---> EEPROM1 : ")
    handleShared.HandleSharedExecCSoI2C0Write(0x54, 0x01, 44)
    bb := handleShared.HandleSharedExecCSoI2C0Read(0x54, 0x01)
	fmt.Println("---> EEPROM2 : ", bb)
	//GoTcpServer测试
	//goTcpServer.GoTcpServer()
	//PCF8563测试
	timeWrite := []int{0,31,11,3,6,12,16}
	handleShared.HandleSharedExecCSoPCFWrite(timeWrite)
	timeBuff := handleShared.HandleSharedExecCSoPCFRead()
	fmt.Println("---> timeBuff: ", timeBuff)

	//GPIO测试
	handleShared.HandleSharedExecCSoGpio37Blinking()

	*/

	//Redis测试
	//goredis.HttpRedisStart()

	/*
	//workerpool测试
	// generate worker to do job
	dispatcher := workerPool.NewDispatcher(3)
	dispatcher.Run()

	// produce job to be done
	producer := workerPool.NewProducer(40)
	producer.Run()

	time.Sleep(time.Second * 2)
	*/
	/*
	//RedisJson测试
	var key01 string
	//var key02 string
	key01 = "Hello02"
	//key02 = "Hello02"
	kvJson := make(map[string]interface{})
	kvJson[key01] = "Redis"
	//kvJson[key02] = "Go"
	handleRedis.HandleRedisJsonInsert(key01, kvJson)
	data := handleRedis.HandleRedisJsonGet(key01)
	fmt.Println("---> Get RedisJson: ", data)

	delv := handleRedis.HandleRedisJsonDel(key01)
	fmt.Println("---> delv: ", delv)
	//handleRedis.HandleRedisJson(key02, kvJson)
	*/
	//MySql测试
	//
	// handleMySql.HandleMySqlPool()
	// handleMySql.HandleDBCreateTable()
	//电参量测试
	
	// var buffElec handleMySql.Buff
	// elecDatas := make([]float64, 5)
	// // elecDatas = append(elecDatas, 1.21, 1.31, 1.401, 1.5001, 1.5)
	
	// for i := 0; i < 5; i++ {
	// 	elecDatas[i] = float64(i)+100.11
	// }
	// fmt.Println("---> elecDatas: ", elecDatas)
	
	
	// handleMySql.HandleDBElecInsert(774, elecDatas, "dbelec")

	// //handleMySql.HandleDBElecDelete(18)
	// fmt.Println("---> Select Single.")
	// buff := buffElec.HandleDBElecGetSingle(2)
	// fmt.Println("---> buff: ", len(buff), buff)
	// fmt.Println("---> Select Manny.")
	// buffElec.HandleDBElecGetManny(1, 10)
	
	//节点测试
	/*
	var buffnode handleMySql.BuffNode
	nodeDatas := make([]int, 10)
	// elecDatas = append(elecDatas, 1.21, 1.31, 1.401, 1.5001, 1.5)
	
	for i := 0; i < 10; i++ {
		nodeDatas[i] = i
	}
	fmt.Println("---> elecDatas: ", nodeDatas)
	
	
	handleMySql.HandleDBNodeInsert(2, nodeDatas, "dbnode")

	handleMySql.HandleDBNodeDelete(1)
	fmt.Println("---> Select Single.")
	buff := buffnode.HandleDBNodeGetSingle(2)
	fmt.Println("---> buff: ", len(buff), buff)
	fmt.Println("---> Select Manny.")
	buffnode.HandleDBNodeGetManny(1, 10)
	*/
	//TenKeys测试
	/*
	var buffTenKeys handleMySql.BuffTenKeys
	TenDatas := make([]int, 10)
	dbname := "dbstagepwm"
	// dbname := "dbspecial"
	// dbname := "dbholiday"

	// elecDatas = append(elecDatas, 1.21, 1.31, 1.401, 1.5001, 1.5)
	
	for i := 0; i < 10; i++ {
		TenDatas[i] = i
	}
	fmt.Println("---> elecDatas: ", TenDatas)
	
	
	handleMySql.HandleDBTenKeysInsert(3, TenDatas, dbname)
	// fmt.Println("---> Delete One.")
	// handleMySql.HandleDBTenKeysDelete(1, dbname)
	fmt.Println("---> Select Single.")
	buff := buffTenKeys.HandleDBTenKeysGetSingle(2, dbname)
	fmt.Println("---> buff: ", len(buff), buff)
	fmt.Println("---> Select Manny.")
	buffTenKeys.HandleDBTenKeysGetManny(1, 10, dbname)
	*/
	//OneKeys测试
	/*
	var buffOneKeys handleMySql.BuffOneKeys
	OneDatas := make([]int, 10)
	dbname := "dbslpwm" //单灯PWM
	// dbname := "dbslgtype" //单灯组类型
	// dbname := "dbslgnum"  //单灯组号
	// dbname := "dbslrelayst" //单灯继电器状态
	// dbname := "dbslhandst"  //单灯手动开关状态

	// elecDatas = append(elecDatas, 1.21, 1.31, 1.401, 1.5001, 1.5)
	
	for i := 0; i < 1; i++ {
		OneDatas[i] = i
	}
	fmt.Println("---> OneDatas: ", OneDatas)
	
	
	handleMySql.HandleDBOneKeysInsert(2, OneDatas, dbname)
	// fmt.Println("---> Delete One.")
	handleMySql.HandleDBOneKeysDelete(1, dbname)
	fmt.Println("---> Select Single.")
	buff := buffOneKeys.HandleDBOneKeysGetSingle(1, dbname)
	fmt.Println("---> buff: ", len(buff), buff)
	fmt.Println("---> Select Manny.")
	buffOneKeys.HandleDBOneKeysGetManny(1, 10, dbname)
	*/
	//单灯定时开关测试
	/*
	var buffTimer handleMySql.BuffTimer
	FourDatas := make([]int, 10)
	dbname := "dbsltimer"

	// elecDatas = append(elecDatas, 1.21, 1.31, 1.401, 1.5001, 1.5)
	
	for i := 0; i < 4; i++ {
		FourDatas[i] = i
	}
	fmt.Println("---> OneDatas: ", FourDatas)
	
	
	handleMySql.HandleDBSlTimerSwitchInsert(3, FourDatas, dbname)
	// fmt.Println("---> Delete One.")
	handleMySql.HandleDBSlTimerSwitchDelete(2)
	fmt.Println("---> Select Single.")
	buff := buffTimer.HandleDBSlTimerSwitchGetSingle(3)
	fmt.Println("---> buff: ", len(buff), buff)
	fmt.Println("---> Select Manny.")
	buffTimer.HandleDBSlTimerSwitchGetManny(1, 10)
	*/
	/*
	//chgoroutines.Main()
	//chgoroutines.SelectMain()
	chgoroutines.TikerDemo()
	chgoroutines.ThreadsPools()

	//godistrabutor.Distrabutor01()
	//godistrabutor.Distrabutor02()
	godistrabutor.Distrabutor03()
	gyjson.GYJson01()
	gyjson.GYJson02()
	gyrwio.GYRWIO01()

	interfaces.Interfaces()
	interfaces.InterfacesNew()

	var ary [3]int
	fmt.Println("---> ", ary[0])
	fmt.Println("---> ", ary[len(ary)-1])

	for i, v := range ary {
		//fmt.Println("---> ", i, v)
		fmt.Printf("---> index[%d] value[%d]\n", i, v)
		ary[i] = 1
	}

	fmt.Println("---> ", ary)
	
	fmt.Println("---> ", reflect.TypeOf(ary))

	var ary01 = [...]string{"1", "2"}
	fmt.Println("---> ", reflect.TypeOf(ary01))

	var a *int = test()
	println(a, *a)

	fmt.Println("---> testOpts")
	testOpts("Hello", ary, ary01, a)

	print("---> Go Redis Base Test <---\n")
	goredis.GoRedisBase()
	print("---> Go Redis Json Test <---\n")
	goredis.GoRedisJson()
	*/
	// gyjson.ChanJson()

	


}

func forRange() {
	sliceData := []int{1, 2, 3, 4, 5, 6}
	for _, v := range sliceData {
		fmt.Println("Value ---> ", v)
	}
}

func test() *int {
	a := 0x100
	return &a
}

func testOpts(s string, any ...interface{}) {
	fmt.Println("---> ", s, any)
}
